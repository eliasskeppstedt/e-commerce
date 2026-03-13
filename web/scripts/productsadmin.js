console.log("productsadmin.js loaded successfully");

document.addEventListener("DOMContentLoaded", () => {
  const adminElement = document.getElementById("adminFlag");
  const isAdmin = adminElement && adminElement.getAttribute("data-is-admin") === "true";
  console.log("Admin script running. isAdmin =", isAdmin);

  // --- Load products from backend ---
  function filterProductsByCategory(categoryId) {
    fetch("/api/products")
      .then(res => res.json())
      .then(products => {
        const grid = document.getElementById("productGrid");
        grid.innerHTML = ""; // clear previous

        const filterId = parseInt(categoryId) || 0;

        products
          .filter(p => !filterId || p.category_id === filterId)
          .forEach(p => renderProduct(p));
      })
      .catch(err => console.error("Failed to load products:", err));
  }

   
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
      </div>
    `;

    // --- Add admin-only edit/delete buttons ---
    if (isAdmin) {
      const contentDiv = card.querySelector(".product-content");
      contentDiv.insertAdjacentHTML("beforeend",
        `<button class="edit-btn">Edit</button><button class="delete-btn">Delete</button>`
      );

      card.querySelector(".edit-btn").onclick = () => enableEditMode(p, card);
      card.querySelector(".delete-btn").onclick = () => deleteProduct(p.product_id);
    }

    grid.appendChild(card);
  }

  // --- Load reviews ---
  function loadReviews(productId) {
    fetch(`/api/reviews?product_id=${productId}`)
      .then(res => res.json())
      .then(reviews => {
        const container = document.querySelector(`.reviews[data-product="${productId}"] .review-list`);
        if (!container) return;
        container.innerHTML = "";

        reviews.forEach(r => {
          const div = document.createElement("div");
          div.className = "review";

          const stars = "★".repeat(r.grade) + "☆".repeat(5 - r.grade);

          let deleteBtnHTML = "";
          if (r.user_id === Number(localStorage.getItem("userId"))) {
            deleteBtnHTML = `<button class="btn-delete-review" data-id="${r.comment_id}" data-product="${productId}">Delete</button>`;
          }

          div.innerHTML = `
            <strong>${r.username}</strong>
            <div class="review-stars">${stars}</div>
            <p>${r.comment_text}</p>
            ${deleteBtnHTML}
          `;

          container.appendChild(div);
        });

        // --- Attach event listeners for review delete buttons ---
        container.querySelectorAll(".btn-delete-review").forEach(btn => {
          btn.onclick = () => {
            const commentId = btn.getAttribute("data-id");
            const productId = btn.getAttribute("data-product");
            deleteReview(commentId, productId);
          };
        });
      })
      .catch(err => console.error("Failed to load reviews:", err));
  }

  // --- Delete a single review ---
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
        alert("Review deleted successfully!");
        loadReviews(productId);
      })
      .catch(err => alert("Error: " + err.message));
  }

  // --- Category filter ---
  const filterSelect = document.getElementById("categorySelect");
  if (filterSelect) {
    filterSelect.addEventListener("change", e => filterProductsByCategory(e.target.value));
  }

  // --- Initial load ---
  filterProductsByCategory("");
});