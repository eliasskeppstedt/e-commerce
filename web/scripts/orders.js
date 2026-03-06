console.log("orders.js loaded successfully");

document.addEventListener("DOMContentLoaded", loadOrders);

function loadOrders() {

  const container = document.getElementById("orders");
  if (!container) return;

  fetch("/api/orders")
    .then(res => res.json())
    .then(orders => {

      container.innerHTML = "";

      if (!orders) return;

      orders.forEach(order => {

        const div = document.createElement("div");

        let itemsHTML = "";

        order.items.forEach(i => {
          itemsHTML += `
            <li>${i.product_name} x${i.quantity} - ${i.price} SEK</li>
          `;
        });

        div.innerHTML = `
          <h3>Order #${order.order_id}</h3>
          <p>Status: ${order.status}</p>
          <p>Date: ${order.date}</p>
          <ul>${itemsHTML}</ul>
        `;

        container.appendChild(div);

      });

    });

}