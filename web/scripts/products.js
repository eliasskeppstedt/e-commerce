console.log("products.js loaded successfully");

const adminElement = document.getElementById("adminFlag");
const isAdmin = adminElement && adminElement.getAttribute("data-is-admin") === "true";

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
  grid.appendChild(card);
}

function filterProductsByCategory(categoryId) {
  fetch("/api/products")
    .then(res => res.json())
    .then(products => {
      const grid = document.getElementById("productGrid");
      grid.innerHTML = "";

      const filterId = parseInt(categoryId) || 0;

      products
        .filter(p => !filterId || p.category_id === filterId)
        .forEach(p => renderProduct(p));
    })
    .catch(err => console.error("Failed to load products:", err));
}

function addToCart(productId) {
  fetch("/api/carts/items", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ product_id: Number(productId) })
  })
    .then(res => {
      if (!res.ok) throw new Error("Failed to add to cart");
      alert("Added to cart!");
      updateCartCount();
    })
    .catch(err => alert(err.message));
}

function updateCartCount() {
  fetch("/api/carts/items")
    .then(res => res.json())
    .then(items => {
      const countElem = document.getElementById("cartCount");
      if (countElem) {
        const totalQty = items.reduce((sum, i) => sum + i.quantity, 0);
        countElem.textContent = totalQty;
      }
    });
}

document.addEventListener("click", function(e) {
  if (e.target.classList.contains("btn-add-to-cart")) {
    const productId = e.target.dataset.id;
    addToCart(productId);
  }
});

// --- Admin: create product ---
function createProduct() {
  if (!isAdmin) return alert("Only admins can add products!");

  const categorySelect = document.getElementById("categorySelectForAdd");
  const categoryId = parseInt(categorySelect.value || "0");
  if (!categoryId) return alert("Please select a category!");

  const stock = parseInt(document.getElementById("stock").value || "0");
  const price = parseFloat(document.getElementById("price").value || "0");

  if (stock < 0) return alert("Stock cannot be negative!");
  if (price < 0) return alert("Price cannot be negative!");

  const product = {
    product_name: document.getElementById("name").value || "",
    manufacturer: document.getElementById("manufacturer").value || "",
    stock: stock,
    price: price,
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

      const currentCategory = document.getElementById("categorySelect").value;
      filterProductsByCategory(currentCategory);
    })
    .catch(err => console.error("Error adding product:", err));
}

document.addEventListener("DOMContentLoaded", () => {
  filterProductsByCategory("");
});
