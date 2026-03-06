document.addEventListener("DOMContentLoaded", loadCart);

function loadCart() {
  fetch("/api/cart")
    .then(res => res.json())
    .then(items => {

      const table = document.getElementById("cart-items");
      table.innerHTML = "";

      if (!items) return;

      let total = 0;

      items.forEach(item => {

        const subtotal = item.price * item.quantity;
        total += subtotal;

        const row = document.createElement("tr");

        row.innerHTML = `
          <td>${item.product_name}</td>
          <td>${item.price} SEK</td>
          <td>${item.quantity}</td>
          <td>${subtotal} SEK</td>
          <td>
            <button class="remove-btn" data-product-id="${item.product_id}">
              ✖
            </button>
          </td>
        `;

        table.appendChild(row);

      });
      document.querySelector(".total-price").innerText = total + " SEK";
    });
}

function renderCart(items) {

  const table = document.getElementById("cart-items");
  table.innerHTML = "";

  let total = 0;

  items.forEach(item => {

    const row = document.createElement("tr");

    const subtotal = item.price * item.quantity;
    total += subtotal;

    row.innerHTML = `
      <td>${item.product_name}</td>
      <td>${item.price} SEK</td>
      <td>${item.quantity}</td>
      <td>${subtotal} SEK</td>
      <td>
        <button class="remove-btn" data-product-id="${item.product_id}">
          ✖
        </button>
      </td>
    `;

    table.appendChild(row);

  });

  document.querySelector(".total-price").innerText = total + " SEK";
}

document.addEventListener("click", e => {

  if (!e.target.classList.contains("remove-btn")) return;

  const productId = e.target.dataset.productId;

  fetch("/api/cart/items/" + productId, {
    method: "DELETE"
  })
  .then(() => loadCart());

});