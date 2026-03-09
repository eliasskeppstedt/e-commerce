console.log("productsadmin.js loaded successfully");

document.addEventListener("DOMContentLoaded", () => {
  const adminElement = document.getElementById("adminFlag");
  const isAdmin = adminElement && adminElement.getAttribute("data-is-admin") === "true";
  console.log("Admin script running. isAdmin =", isAdmin);

  // --- Render a product card ---
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

    // Admin-only buttons
    const contentDiv = card.querySelector(".product-content");
    contentDiv.insertAdjacentHTML("beforeend",
      `<button class="edit-btn">Edit</button><button class="delete-btn">Delete</button>`
    );

    card.querySelector(".edit-btn").onclick = () => enableEditMode(p, card);
    card.querySelector(".delete-btn").onclick = () => deleteProduct(p.product_id);

    grid.appendChild(card);
  }

  // --- Load products ---
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

  // --- Add new product ---
  function addProduct() {
    if (!isAdmin) return alert("Only admins can add products!");

    const categorySelect = document.getElementById("categorySelectForAdd");
    const categoryId = parseInt(categorySelect.value || "0");
    if (!categoryId) return alert("Please select a category!");

    let stock = parseInt(document.getElementById("stock").value || "0");
    let price = parseFloat(document.getElementById("price").value || "0");

    // Prevent negative stock/price
    if (stock < 0 || price < 0) return alert("Stock and price cannot be negative!");

    const product = {
      product_name: document.getElementById("name").value,
      manufacturer: document.getElementById("manufacturer").value || "",
      stock,
      price,
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

  // --- Enable edit mode ---
  function enableEditMode(product, card) {
    const stockSpan = card.querySelector(".stock");
    const priceSpan = card.querySelector(".price");

    const stockInput = document.createElement("input");
    stockInput.type = "number";
    stockInput.className = "edit-stock";
    stockInput.value = product.stock;
    stockInput.min = "0";

    const priceInput = document.createElement("input");
    priceInput.type = "number";
    priceInput.className = "edit-price";
    priceInput.step = "0.01";
    priceInput.value = product.price;
    priceInput.min = "0";

    // Prevent typing negative values
    stockInput.addEventListener("input", () => {
      if (parseInt(stockInput.value) < 0) stockInput.value = 0;
    });
    priceInput.addEventListener("input", () => {
      if (parseFloat(priceInput.value) < 0) priceInput.value = 0;
    });

    stockSpan.innerHTML = "";
    stockSpan.appendChild(stockInput);
    priceSpan.innerHTML = "";
    priceSpan.appendChild(priceInput);

    const editBtn = card.querySelector(".edit-btn");
    editBtn.textContent = "Save";
    editBtn.onclick = () => saveUpdate(product.product_id, card);
  }

  // --- Save updated product ---
  function saveUpdate(productId, card) {
    const newStock = parseInt(card.querySelector(".edit-stock").value);
    const newPrice = parseFloat(card.querySelector(".edit-price").value);

    // Prevent negative values
    if (newStock < 0 || newPrice < 0) return alert("Stock and price cannot be negative!");

    fetch(`/api/products/${productId}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ stock: newStock, price: newPrice })
    })
      .then(res => res.json())
      .then(() => {
        const currentCategory = document.getElementById("categorySelect").value;
        filterProductsByCategory(currentCategory);
      })
      .catch(err => console.error("Update failed:", err));
  }

  // --- Delete product ---
  function deleteProduct(productId) {
    if (!confirm("Are you sure you want to delete this product?")) return;

    fetch(`/api/products/${productId}`, { method: "DELETE" })
      .then(res => res.json())
      .then(() => {
        const currentCategory = document.getElementById("categorySelect").value;
        filterProductsByCategory(currentCategory);
      })
      .catch(err => console.error("Delete failed:", err));
  }

  // --- Add product button ---
  const addBtn = document.querySelector(".add-product button");
  if (addBtn) addBtn.onclick = addProduct;

  // --- Category filter ---
  const filterSelect = document.getElementById("categorySelect");
  if (filterSelect) {
    filterSelect.addEventListener("change", e => {
      filterProductsByCategory(e.target.value);
    });
  }

  // --- Initial load ---
  filterProductsByCategory("");
});
