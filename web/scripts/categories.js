console.log("categories.js loaded successfully");

function loadCategories() {
  fetch("/api/categories")
    .then(res => res.json())
    .then(categories => {
      const filterSelect = document.getElementById("categorySelect");
      if (filterSelect) {
        filterSelect.innerHTML = '<option value="">All Categories</option>';
        categories.forEach(c => {
          const option = document.createElement("option");
          option.value = c.category_id;
          option.textContent = c.category_name;
          filterSelect.appendChild(option);
        });
      }

      const addSelect = document.getElementById("categorySelectForAdd");
      if (addSelect) {
        addSelect.innerHTML = '<option value="">Select a Category</option>';
        categories.forEach(c => {
          const option = document.createElement("option");
          option.value = c.category_id;
          option.textContent = c.category_name;
          addSelect.appendChild(option);
        });
      }
    })
    .catch(err => console.error("Failed to load categories:", err));
}

// --- Filter products when dropdown changes ---
const filterSelect = document.getElementById("categorySelect");
if (filterSelect) {
  filterSelect.addEventListener("change", (e) => {
    filterProductsByCategory(e.target.value);
  });
}

// --- Admin add category ---
const addBtn = document.getElementById("addCategoryBtn");
const newInput = document.getElementById("newCategoryInput");
if (addBtn && newInput) {
  addBtn.addEventListener("click", () => {
    const newCategory = newInput.value.trim();
    if (!newCategory) return alert("Please enter a category name.");

    fetch("/api/categories", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ category_name: newCategory })
    })
      .then(() => {
        newInput.value = "";
        loadCategories();
      })
      .catch(err => console.error("Failed to add category:", err));
  });
}

// --- Admin remove category ---
const removeBtn = document.getElementById("removeCategoryBtn");
if (removeBtn) {
  removeBtn.addEventListener("click", () => {
    const select = document.getElementById("categorySelect");
    const categoryId = select?.value;
    if (!categoryId) return alert("Please select a category to delete.");
    if (!confirm("Delete this category? This will remove it from DB.")) return;

    fetch(`/api/categories/${categoryId}`, { method: "DELETE" })
      .then(res => res.json())
      .then(data => {
        if (data.error) {
          alert(data.error);
          return;
        }
        loadCategories();
        filterProductsByCategory(document.getElementById("categorySelect")?.value || "");
      })
      .catch(err => console.error("Failed to delete category:", err));
  });
}

// --- Initial load ---
loadCategories();
