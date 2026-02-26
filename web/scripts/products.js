console.log("products.js loaded successfully");

// --- Render a single product card ---
function renderProduct(p) {
  const grid = document.getElementById("productGrid");
  const card = document.createElement("div");
  card.className = "product-card";

  card.innerHTML = `
    <div class="product-image"></div>
    <div class="product-content">
      <h3>${p.product_name}</h3>
      <p>${p.description}</p>
      <p>Manufacturer: ${p.manufacturer}</p>
      <p>Stock: ${p.stock}</p>
      <p>Category: ${p.category_name}</p>
      <strong>${p.price} SEK</strong>
      <br />
      <button class="delete-btn">Delete</button>
    </div>
  `;

  card.querySelector(".delete-btn").onclick = () =>
    deleteProduct(p.product_id);

  grid.appendChild(card);
}

// --- Load and filter products ---
function filterProductsByCategory(categoryId) {
  fetch("/api/products")
    .then(res => res.json())
    .then(products => {
      const grid = document.getElementById("productGrid");
      grid.innerHTML = ""; // Clear previous products

      // Parse categoryId to int
      const filterId = parseInt(categoryId) || 0;

      products
        .filter(p => !filterId || p.category_id === filterId)
        .forEach(p => renderProduct(p));
    })
    .catch(err => console.error("Failed to load products:", err));
}

// --- Add new product ---
function addProduct() {
  const categorySelect = document.getElementById("categorySelectForAdd");
  const categoryId = parseInt(categorySelect.value || "0");
  if (!categoryId) return alert("Please select a category!");

  const product = {
    product_name: document.getElementById("name").value,
    manufacturer: document.getElementById("manufacturer").value || "",
    stock: parseInt(document.getElementById("stock").value || "0"),
    price: parseFloat(document.getElementById("price").value || "0"),
    description: document.getElementById("description").value || "",
    category_id: categoryId
  };

  fetch("/api/products", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(product)
  })
    .then(res => res.json())
    .then(() => {
      // Clear form
      document.getElementById("name").value = "";
      document.getElementById("manufacturer").value = "";
      document.getElementById("stock").value = "";
      document.getElementById("price").value = "";
      document.getElementById("description").value = "";
      categorySelect.value = "";

      // Refresh products with current filter
      const currentCategory = document.getElementById("categorySelect").value;
      filterProductsByCategory(currentCategory);
    })
    .catch(err => console.error("Error adding product:", err));
}

// --- Delete product ---
function deleteProduct(productId) {
  if (!confirm("Are you sure you want to delete this product?")) return;

  fetch(`/api/products/${productId}`, { method: "DELETE" })
    .then(res => res.json())
    .then(() => {
      const currentCategory = document.getElementById("categorySelect").value;
      filterProductsByCategory(currentCategory);
    })
    .catch(err => console.error("Delete failed:", err));
}

// --- Initial load ---
filterProductsByCategory(""); // load all products initially
