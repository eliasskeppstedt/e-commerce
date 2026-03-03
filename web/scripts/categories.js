console.log("categories.js loaded successfully");

// --- Laddar in kategorierna från backend ---
function loadCategories() {
  fetch("/api/categories")
    .then(res => res.json())
    .then(categories => {
      // --- Dropdown filter ---
      const filterSelect = document.getElementById("categorySelect");
      filterSelect.innerHTML = '<option value="">All Categories</option>';
      categories.forEach(c => {
        const option = document.createElement("option");
        option.value = c.category_id;      // use numeric ID
        option.textContent = c.category_name;
        filterSelect.appendChild(option);
      });

      // --- Lägger till kategori till produkt menyn (där man väljer kategori) ---
      const addSelect = document.getElementById("categorySelectForAdd");
      addSelect.innerHTML = '<option value="">Select a Category</option>';
      categories.forEach(c => {
        const option = document.createElement("option");
        option.value = c.category_id;      // numeric ID
        option.textContent = c.category_name;
        addSelect.appendChild(option);
      });
    })
    .catch(err => console.error("Failed to load categories:", err));
}

// --- Dropdown meny uppdateras ---
document.getElementById("categorySelect").addEventListener("change", (e) => {
  filterProductsByCategory(e.target.value);
});

// --- Lägger till kategori --- Måste fixa så bara admin kan göra detta
document.getElementById("addCategoryBtn").addEventListener("click", () => {
  const newCategory = document.getElementById("newCategoryInput").value.trim();
  if (!newCategory) return alert("Please enter a category name.");

  fetch("/api/categories", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ category_name: newCategory })
  })
    .then(res => res.json())
    .then(() => {
      document.getElementById("newCategoryInput").value = "";
      loadCategories(); // reload dropdowns
    })
    .catch(err => console.error("Failed to add category:", err));
});

// --- Tar bort kategori --- måste fixa så bara admin kan göra detta
document.getElementById("removeCategoryBtn").addEventListener("click", () => {
  const select = document.getElementById("categorySelect");
  const categoryId = select.value;
  if (!categoryId) return alert("Please select a category to delete.");

  if (!confirm("Delete this category? This will remove it from DB.")) return;

  fetch(`/api/categories/${categoryId}`, { method: "DELETE" })
    .then(res => res.json())
    .then(data => {
      if (data.error) {
        alert(data.error);
        return;
      }
      // reload dropdowns and refresh products for current filter
      loadCategories();
      filterProductsByCategory(document.getElementById("categorySelect").value);
    })
    .catch(err => console.error("Failed to delete category:", err));
});

// --- Initial load ---
loadCategories();
