console.log("products.js loaded successfully");

const adminElement = document.getElementById("adminFlag");
const isAdmin = adminElement && adminElement.getAttribute("data-is-admin") === "true";
const currentUserId = adminElement ? parseInt(adminElement.getAttribute("data-user-id")) : null;

// -------------------------
// Render product
// -------------------------
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

      <button class="btn-add-to-cart" data-id="${p.product_id}">
        Add to cart
      </button>

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

        <button class="btn-add-review" data-id="${p.product_id}">
          Submit Review
        </button>
      </div>
    </div>
  `;

  grid.appendChild(card);

  loadReviews(p.product_id);
}


// -------------------------
// Load products
// -------------------------
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


// -------------------------
// Add to cart
// -------------------------
function addToCart(productId) {

  fetch("/api/carts/items", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      product_id: Number(productId)
    })
  })
    .then(res => {
      if (!res.ok) throw new Error("Failed to add to cart");
      alert("Added to cart!");
      updateCartCount();
    })
    .catch(err => alert(err.message));
}


// -------------------------
// Update cart counter
// -------------------------
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


// -------------------------
// Load reviews
// -------------------------
function loadReviews(productId) {

  fetch(`/api/reviews?product_id=${productId}`)
    .then(res => res.json())
    .then(reviews => {

      const container = document.querySelector(
        `.reviews[data-product="${productId}"] .review-list`
      );

      if (!container) return;

      container.innerHTML = "";

      reviews.forEach(r => {

        const div = document.createElement("div");
        div.className = "review";
        div.dataset.userId = r.user_id;

        const stars = "★".repeat(r.grade) + "☆".repeat(5 - r.grade);

        let deleteBtnHtml = "";
        if (r.user_id === currentUserId) {
          deleteBtnHtml = `<button class="btn-delete-review" data-id="${r.comment_id}">Delete</button>`;
        }

        div.innerHTML = `
          <strong>${r.username}</strong>
          <div class="review-stars">${stars}</div>
          <p>${r.comment_text}</p>
          ${deleteBtnHtml}
        `;

        container.appendChild(div);
      });

    })
    .catch(err => console.error("Failed to load reviews:", err));
}


// -------------------------
// Add review
// -------------------------
function addReview(productId, text, grade) {

  fetch("/api/reviews", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      product_id: Number(productId),
      text: text,
      grade: Number(grade)
    })
  })
    .then(res => {
      if (!res.ok) throw new Error("Failed to add review");
  
      loadReviews(productId);
    })
    .catch(err => alert(err.message));
}


// -------------------------
// Delete review
// -------------------------
function deleteReview(commentId, productId) {

  fetch("/api/reviews", {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ comment_id: Number(commentId) })
  })
    .then(res => {
      if (!res.ok) throw new Error("Failed to delete review");
  
      loadReviews(productId);
    })
    .catch(err => alert(err.message));
}


// -------------------------
// Global click handler
// -------------------------
document.addEventListener("click", function(e) {

  // Add to cart
  if (e.target.classList.contains("btn-add-to-cart")) {

    const productId = e.target.dataset.id;
    addToCart(productId);

  }

  // Add review
  if (e.target.classList.contains("btn-add-review")) {

    const productId = e.target.dataset.id;

    const container = e.target.closest(".reviews");

    const textarea = container.querySelector(".review-text");
    const gradeSelect = container.querySelector(".review-grade");

    const text = textarea.value.trim();
    const grade = gradeSelect.value;

    if (!grade) return alert("Please select a rating (1–5 stars)");
    if (!text) return alert("Write a review first!");

    addReview(productId, text, grade);

    textarea.value = "";
    gradeSelect.value = "";
  }

  // Delete review
  if (e.target.classList.contains("btn-delete-review")) {
    const commentId = e.target.dataset.id;
    const productId = e.target.closest(".reviews").dataset.product;
    deleteReview(commentId, productId);
  }

});


// -------------------------
// Admin: create product
// -------------------------
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
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(product)
  })
    .then(res => res.json())
    .then(() => {

      document.getElementById("name").value = "";
      document.getElementById("manufacturer").value = "";
      document.getElementById("stock").value = "";
      document.getElementById("price").value = "";
      document.getElementById("description").value = "";

      categorySelect.value = "";

      const currentCategory =
        document.getElementById("categorySelect").value;

      filterProductsByCategory(currentCategory);
    })
    .catch(err => console.error("Error adding product:", err));
}


// -------------------------
// Initial load
// -------------------------
document.addEventListener("DOMContentLoaded", () => {

  filterProductsByCategory("");
  updateCartCount();

});