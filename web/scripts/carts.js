console.log("carts.js loaded successfully");

const cartTableBody = document.querySelector(".cart-table tbody");
const totalPriceElem = document.querySelector(".total-price");

// --- Fetch and render cart items ---
function loadCart() {
  fetch("/api/carts/items")
    .then(res => res.json())
    .then(items => {
      cartTableBody.innerHTML = ""; // clear table
      let total = 0;

      items.forEach(item => {
        const row = document.createElement("tr");

        row.innerHTML = `
          <td>${item.product_name}</td>
          <td>${item.product_price} SEK</td>
          <td>${item.quantity}</td>
          <td>${item.subtotal.toFixed(2)} SEK</td>
          <td><button class="remove-btn" data-id="${item.product_id}">✖</button></td>
        `;
        cartTableBody.appendChild(row);

        total += item.subtotal;
      });

      totalPriceElem.textContent = total.toFixed(2);
    })
    .catch(err => console.error("Failed to load cart items:", err));
}

// --- Remove item from cart ---
function removeFromCart(productId) {
  fetch("/api/carts/items", {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ product_id: Number(productId) })
  })
    .then(res => {
      if (!res.ok) throw new Error("Failed to remove item");
      loadCart(); // reload cart after removal
    })
    .catch(err => alert(err.message));
}

// --- Checkout ---
function checkout() {
  fetch("/api/carts/checkout", { method: "POST" })
    .then(res => {
      if (!res.ok) throw new Error("Checkout failed");
      alert("Purchase successful!");
      loadCart(); // reload cart, should be empty
    })
    .catch(err => alert(err.message));
}

// --- Event delegation for remove buttons and checkout ---
document.addEventListener("click", function(e) {
  if (e.target.classList.contains("remove-btn")) {
    const productId = e.target.dataset.id;
    removeFromCart(productId);
  } else if (e.target.classList.contains("checkout-btn")) {
    checkout();
  }
});

// --- Initial load ---
document.addEventListener("DOMContentLoaded", loadCart);
