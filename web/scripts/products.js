console.log("products.js loaded successfully");

// --- Admin flag from hidden HTML ---
const adminElement = document.getElementById("adminFlag");
const isAdmin = adminElement && adminElement.getAttribute("data-is-admin") === "true";

// --- Render a single product (no edit/delete buttons) ---
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
      <p>Stock: <span class="stock">${p.stock}</span></p>
      <p>Price: <span class="price">${p.price}</span> SEK</p>
      <p>Category: ${p.category_name}</p>
      <button class="btn-add-to-cart" data-id="${p.product_id}">Add to cart</button>
    </div>
  `;
  console.log(p)
  grid.appendChild(card);
}

// --- Filter products by category ---
function filterProductsByCategory(categoryId) {
  fetch("/api/products")
    .then(res => res.json())
    .then(products => {
      const grid = document.getElementById("productGrid");
      grid.innerHTML = ""; // Clear previous

      const filterId = parseInt(categoryId) || 0;

      products
        .filter(p => !filterId || p.category_id === filterId)
        .forEach(p => renderProduct(p));
    })
    .catch(err => console.error("Failed to load products:", err));
}

document.addEventListener("click", function(e) {
  if (e.target.classList.contains("btn-add-to-cart")) {
    const productId = e.target.dataset.id;
    addToCart(productId);
  }
});

function addToCart(productId) {
  console.log("Adding product:", productId, "type of productId:", typeof productId);

  fetch("/api/cart/items", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      product_id: Number(productId)
    })
  });
}

// --- Add a new product (admin only) ---
function createProduct() {
  if (!isAdmin) return alert("Only admins can add products!");

  const categorySelect = document.getElementById("categorySelectForAdd");
  const categoryId = parseInt(categorySelect.value || "0");
  if (!categoryId) return alert("Please select a category!");

  const product = {
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

      // Refresh products
      const currentCategory = document.getElementById("categorySelect").value;
      filterProductsByCategory(currentCategory);
    })
    .catch(err => console.error("Error adding product:", err));
}

// --- Initial load ---
document.addEventListener("DOMContentLoaded", () => {
  filterProductsByCategory("");
});
