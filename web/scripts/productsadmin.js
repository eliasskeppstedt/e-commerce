console.log("productadmin.js loaded successfully");

// -------------------------
// User / Admin info
// -------------------------
const adminElement = document.getElementById("adminFlag");
const isAdmin = adminElement?.dataset.isAdmin === "true" || false;
const userId = adminElement?.dataset.userId
  ? Number(adminElement.dataset.userId)
  : null;

// -------------------------
// Render a single product
// -------------------------
function renderProduct(p) {
  const grid = document.getElementById("productGrid");
  if (!grid) return;

  const card = document.createElement("div");
  card.className = "product-card";

  const productName = p.product_name || "Unnamed Product";
  const description = p.description || "";
  const manufacturer = p.manufacturer || "Unknown";
  const stock = p.stock != null ? p.stock : 0;
  const price = p.price != null ? p.price.toFixed(2) : "0.00";
  const categoryName = p.category_name || "Uncategorized";

  card.innerHTML = `
    <div class="product-image"></div>
    <div class="product-content">
      <h3>${productName}</h3>
      <p>${description}</p>
      <p>Manufacturer: ${manufacturer}</p>
      <p>Stock: <span class="stock">${stock}</span></p>
      <p>Price: <span class="price">${price}</span> SEK</p>
      <p>Category: ${categoryName}</p>

      <button class="btn-add-to-cart" data-id="${p.product_id}">Add to cart</button>
      ${isAdmin ? `<button class="btn-delete-product" data-id="${p.product_id}">Remove product</button>` : ``}

      <div class="reviews" data-product="${p.product_id}">
        <h4>Reviews</h4>
        <div class="review-list"></div>
        <select class="review-grade">
          <option value="">Rating</option>
          <option value="5">★★★★★ (5)</option>
          <option value="4">★★★★ (4)</option>
          <option value="3">★★★ (3)</option>
          <option value="2">★★ (2)</option>
          <option value="1">★ (1)</option>
        </select>
        <textarea class="review-text" placeholder="Write a review"></textarea>
        <button class="btn-add-review" data-id="${p.product_id}">Submit Review</button>
      </div>
    </div>
  `;

  // Admin: add edit button
  if (isAdmin) {
    const contentDiv = card.querySelector(".product-content");
    const editBtn = document.createElement("button");
    editBtn.textContent = "Edit";
    editBtn.className = "edit-btn";
    editBtn.onclick = () => enableEditMode(p, card);
    contentDiv.appendChild(editBtn);
  }

  grid.appendChild(card);
  loadReviews(p.product_id);
}

// -------------------------
// Load products with optional category filter
// -------------------------
function filterProductsByCategory(categoryId) {
  fetch("/api/products")
    .then(res => res.json())
    .then(products => {
      if (!Array.isArray(products)) products = [];

      const grid = document.getElementById("productGrid");
      if (!grid) return;
      grid.innerHTML = "";

      const filterId = parseInt(categoryId) || 0;

      products
        .filter(p => !filterId || p.category_id === filterId)
        .forEach(p => renderProduct(p));
    })
    .catch(err => console.error("Failed to load products:", err));
}

// -------------------------
// Add to cart
// -------------------------
function addToCart(productId) {
  fetch("/api/carts/items", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ product_id: Number(productId) })
  })
    .then(async res => {
      if (!res.ok) {
        const data = await res.json().catch(() => ({}));
        throw new Error(data.error || "Failed to add to cart");
      }
      alert("Added to cart!");
      updateCartCount();
    })
    .catch(err => alert("Error: " + err.message));
}

// -------------------------
// Update cart counter
// -------------------------
function updateCartCount() {
  fetch("/api/carts/items")
    .then(res => res.json())
    .then(items => {
      const countElem = document.getElementById("cartCount");
      if (!countElem) return;
      const totalQty = Array.isArray(items) ? items.reduce((sum, i) => sum + (i.quantity || 0), 0) : 0;
      countElem.textContent = totalQty;
    })
    .catch(err => console.error("Failed to update cart count:", err));
}

// -------------------------
// Load reviews
// -------------------------
function loadReviews(productId) {
  fetch(`/api/reviews?product_id=${productId}`)
    .then(res => res.json())
    .then(reviews => {
      const container = document.querySelector(`.reviews[data-product="${productId}"] .review-list`);
      if (!container) return;
      container.innerHTML = "";

      if (!Array.isArray(reviews)) reviews = [];

      reviews.forEach(r => {
        const div = document.createElement("div");
        div.className = "review";
        const stars = "★".repeat(r.grade || 0) + "☆".repeat(5 - (r.grade || 0));

        let deleteBtnHTML = "";
        if (isAdmin || r.user_id === userId) {
          deleteBtnHTML = `<button class="btn-delete-review" data-id="${r.comment_id}" data-product="${productId}">Delete</button>`;
        }

        div.innerHTML = `
          <strong>${r.username || "Anonymous"}</strong>
          <div class="review-stars">${stars}</div>
          <p>${r.comment_text || ""}</p>
          ${deleteBtnHTML}
        `;

        container.appendChild(div);
      });
    })
    .catch(err => console.error("Failed to load reviews:", err));
}

// -------------------------
// Add review
// -------------------------
function addReview(productId, text, grade, createdAt) {
  fetch("/api/reviews", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      product_id: Number(productId),
      text: text,
      grade: Number(grade),
      created_at: createdAt
    })
  })
    .then(async res => {
      if (!res.ok) {
        const data = await res.json().catch(() => ({}));
        throw new Error(data.error || "Failed to add review");
      }
      alert("Review added successfully!");
      loadReviews(productId);
    })
    .catch(err => alert("Error: " + err.message));
}

// -------------------------
// Delete review
// -------------------------
function deleteReview(commentId, productId) {
  fetch("/api/reviews", {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ comment_id: Number(commentId) })
  })
    .then(async res => {
      if (!res.ok) {
        const data = await res.json().catch(() => ({}));
        throw new Error(data.error || "Failed to delete review");
      }
      loadReviews(productId);
    })
    .catch(err => alert("Error: " + err.message));
}

// -------------------------
// Admin: edit product
// -------------------------
function enableEditMode(product, card) {
  const stockSpan = card.querySelector(".stock");
  const priceSpan = card.querySelector(".price");

  stockSpan.innerHTML = `<input type="number" class="edit-stock" value="${product.stock}">`;
  priceSpan.innerHTML = `<input type="number" step="0.01" class="edit-price" value="${product.price}">`;

  const editBtn = card.querySelector(".edit-btn");
  editBtn.textContent = "Save";
  editBtn.onclick = () => saveUpdate(product.product_id, card);
}

function saveUpdate(productId, card) {
  const newStock = parseInt(card.querySelector(".edit-stock").value);
  const newPrice = parseFloat(card.querySelector(".edit-price").value);

  if (newStock < 0 || newPrice < 0) return alert("Stock and Price must be >= 0");

  fetch(`/api/products/${productId}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ stock: newStock, price: newPrice })
  })
    .then(async res => {
      if (!res.ok) {
        const data = await res.json().catch(() => ({}));
        throw new Error(data.error || "Update failed");
      }
      alert("Product updated successfully!");
      filterProductsByCategory(document.getElementById("categorySelect")?.value || "");
    })
    .catch(err => alert("Error: " + err.message));
}

// -------------------------
// Admin: create product
// -------------------------
function createProduct() {
  if (!isAdmin) return alert("Only admins can add products!");

  const categorySelect = document.getElementById("categorySelectForAdd");
  const categoryId = parseInt(categorySelect?.value || "0");
  if (!categoryId) return alert("Please select a category!");

  const stock = parseInt(document.getElementById("stock")?.value || "0");
  const price = parseFloat(document.getElementById("price")?.value || "0");

  if (stock < 0 || price < 0) return alert("Stock and Price must be >= 0");

  const product = {
    product_name: document.getElementById("name")?.value || "",
    manufacturer: document.getElementById("manufacturer")?.value || "",
    stock,
    price,
    description: document.getElementById("description")?.value || "",
    category_id: categoryId
  };

  fetch("/api/products", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(product)
  })
    .then(async res => {
      if (!res.ok) {
        const data = await res.json().catch(() => ({}));
        throw new Error(data.error || "Failed to add product");
      }
      alert("Product added successfully!");
      ["name","manufacturer","stock","price","description"].forEach(id => {
        const elem = document.getElementById(id);
        if (elem) elem.value = "";
      });
      if (categorySelect) categorySelect.value = "";
      filterProductsByCategory(document.getElementById("categorySelect")?.value || "");
    })
    .catch(err => alert("Error: " + err.message));
}

// -------------------------
// Delete product
// -------------------------
function deleteProduct(productId) {
  if (!confirm("Are you sure you want to delete this product?")) return;

  fetch(`/api/products/${productId}`, { method: "DELETE" })
    .then(res => res.json())
    .then(() => filterProductsByCategory(document.getElementById("categorySelect")?.value || ""))
    .catch(err => console.error("Delete failed:", err));
}

// -------------------------
// Global click handler
// -------------------------
document.addEventListener("click", e => {
  if (e.target.classList.contains("btn-add-to-cart")) addToCart(e.target.dataset.id);
  if (e.target.classList.contains("btn-delete-product")) deleteProduct(e.target.dataset.id);

  if (e.target.classList.contains("btn-add-review")) {
    const container = e.target.closest(".reviews");
    const productId = e.target.dataset.id;
    const text = container.querySelector(".review-text")?.value.trim();
    const grade = container.querySelector(".review-grade")?.value;
    if (!grade) return alert("Please select a rating (1–5 stars)");
    if (!text) return alert("Write a review first!");

    const createdAt = new Date().toISOString();
    addReview(productId, text, grade, createdAt);

    container.querySelector(".review-text").value = "";
    container.querySelector(".review-grade").value = "";
  }

  if (e.target.classList.contains("btn-delete-review")) {
    deleteReview(e.target.dataset.id, e.target.dataset.product);
  }
});

// -------------------------
// Initial load
// -------------------------
document.addEventListener("DOMContentLoaded", () => {
  filterProductsByCategory("");
  updateCartCount();
});
