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

<h3 class="product-title">${p.product_name}</h3>

<div class="product-main">

  <div class="product-left">
    <p>${p.description}</p>
    <p>Manufacturer: ${p.manufacturer}</p>
    <p>Stock: <span class="stock">${p.stock}</span></p>
    <p>Price: <span class="price">${p.price}</span> SEK</p>
    <p>Category: ${p.category_name}</p>

    <button class="add-to-cart"
            data-product-id="${p.product_id}"
            ${p.stock === 0 ? "disabled" : ""}>
      ${p.stock === 0 ? "Out of stock" : "Add to cart"}
    </button>
  </div>

  <div class="product-right">

    <h4>Write Review</h4>

    <textarea
      placeholder="Write a review..."
      class="review-text"
      data-product-id="${p.product_id}">
    </textarea>

    <select class="review-grade" data-product-id="${p.product_id}">
      <option value="1">1⭐</option>
      <option value="2">2⭐</option>
      <option value="3">3⭐</option>
      <option value="4">4⭐</option>
      <option value="5">5⭐</option>
    </select>

    <button class="add-review" data-product-id="${p.product_id}">
      Submit Review
    </button>

  </div>

</div>

<div class="reviews">

  <h4>Reviews</h4>

  <div class="review-list" id="reviews-${p.product_id}"></div>

</div>
`;
  grid.appendChild(card);
  loadReviews(p.product_id);
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

// --- Add a new product (admin only) ---
function addProduct() {
  if (!isAdmin) return alert("Only admins can add products!");

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

document.addEventListener("click", (e) => {

  if (!e.target.classList.contains("add-to-cart")) return;

  console.log("button clicked", e.target);

  const productId = e.target.dataset.productId;

  console.log("productId:", productId);

  fetch("/api/product/add", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      product_id: parseInt(productId)
    })
  })
    .then(res => {
      if (!res.ok) throw new Error("Out of stock");
      return res.json();
    })
    .then(() => {

      const currentCategory =
        document.getElementById("categorySelect").value;

      filterProductsByCategory(currentCategory);

    })
    .catch(err => {
      alert(err.message);
    });

});

document.addEventListener("click", e => {

  if (!e.target.classList.contains("add-review")) return;

  const productId = e.target.dataset.productId;

  const text = document.querySelector(
    `.review-text[data-product-id="${productId}"]`
  ).value;

  const grade = document.querySelector(
    `.review-grade[data-product-id="${productId}"]`
  ).value;

  fetch("/api/reviews", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      product_id: parseInt(productId),
      text: text,
      grade: parseInt(grade)
    })
  })
    .then(res => res.json())
    .then(() => {
      loadReviews(productId);
    });

});

function loadReviews(productId) {

  fetch("/api/reviews/" + productId)
    .then(res => res.json())
    .then(reviews => {

      const container = document.getElementById("reviews-" + productId);
      container.innerHTML = "";

      if (!reviews) return;

      reviews.forEach(r => {

        const div = document.createElement("div");

        div.innerHTML = `
          <p>⭐ ${r.grade}/5</p>
          <p>${r.comment_text}</p>

          <button class="delete-review" 
                  data-review-id="${r.comment_id}"
                  data-product-id="${productId}">
            Delete
          </button>
        `;

        container.appendChild(div);

      });

    });
}
document.addEventListener("click", e => {

  if (!e.target.classList.contains("delete-review")) return;

  const reviewId = e.target.dataset.reviewId;
  const productId = e.target.dataset.productId;

  fetch("/api/reviews/" + reviewId, {
    method: "DELETE"
  })
    .then(() => {
      loadReviews(productId);
    });

});