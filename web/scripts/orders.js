console.log("orders.js loaded");
document.addEventListener("DOMContentLoaded", () => {
  loadOrders();
});

const adminElement = document.getElementById("adminFlag");
const isAdmin = adminElement && adminElement.getAttribute("data-is-admin") === "true";

function loadOrders() {

  const endpoint = isAdmin ? "/api/admin/orders" : "/api/orders";

  fetch(endpoint)
    .then(res => res.json())
    .then(data => {

      const container = document.getElementById("ordersContainer");
      container.innerHTML = "";

      if (isAdmin) {
        renderAdminOrders(data, container);
      } else {
        renderUserOrders(data, container);
      }
    })
    .catch(err => console.error("Failed to load orders", err));
}

function renderUserOrders(orders, container) {

  orders.forEach(order => {

    const orderDiv = document.createElement("div");
    orderDiv.className = "order-card";

    orderDiv.innerHTML = `
      <h3>Order #${order.order_id}</h3>
      <p>Status: ${order.status}</p>
      <p>Date: ${order.date}</p>
      <div class="order-items"></div>
    `;

    const itemsContainer = orderDiv.querySelector(".order-items");

    order.items.forEach(item => {

      const itemDiv = document.createElement("div");
      itemDiv.className = "order-item";

      itemDiv.innerHTML = `
        ${item.product_name}  
        x${item.quantity}  
        ${item.price} SEK
      `;

      itemsContainer.appendChild(itemDiv);
    });

    container.appendChild(orderDiv);
  });
}

function renderAdminOrders(users, container) {

  users.forEach(user => {

    const userDiv = document.createElement("div");
    userDiv.className = "user-orders";

    userDiv.innerHTML = `
      <h2>User: ${user.username} (ID: ${user.user_id})</h2>
      <div class="user-orders-list"></div>
    `;

    const ordersContainer = userDiv.querySelector(".user-orders-list");

    user.orders.forEach(order => {

      const orderDiv = document.createElement("div");
      orderDiv.className = "order-card";

      orderDiv.innerHTML = `
        <h3>Order #${order.order_id}</h3>
        <p>Status: ${order.status}</p>
        <p>Date: ${order.date}</p>
        <div class="order-items"></div>
      `;

      const itemsContainer = orderDiv.querySelector(".order-items");

      order.items.forEach(item => {

        const itemDiv = document.createElement("div");
        itemDiv.className = "order-item";

        itemDiv.innerHTML = `
          ${item.product_name}  
          x${item.quantity}  
          ${item.price} SEK
        `;

        itemsContainer.appendChild(itemDiv);
      });

      ordersContainer.appendChild(orderDiv);
    });

    container.appendChild(userDiv);
  });
}