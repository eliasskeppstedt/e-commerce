console.log("products.js loaded successfully");

// Visar ett produkt "kort"
function renderProduct(p) {
  const grid = document.getElementById("productGrid");
  const card = document.createElement("div");
  card.className = "product-card";

  card.innerHTML = `
    <div class="product-image"></div>
    <div class="product-content">
      <h3>${p.product_name}</h3>
      <p>Description: ${p.description}</p>
      <p>Manufacturer: ${p.manufacturer}</p>
      <p>Stock: ${p.stock}</p>
      <p>Category: ${p.category_name}</p>
      <strong>${p.price} SEK</strong>
    </div>
  `;

  grid.appendChild(card);
}

// Laddar in produkter från servern och renderar dom
fetch("/api/products")
  .then(res => res.json())
  .then(products => {
    const grid = document.getElementById("productGrid");
    grid.innerHTML = ""; // Clear any placeholder content
    products.forEach(p => renderProduct(p));
  })
  .catch(err => console.error("Failed to load products:", err));

// Lägger till en ny produkt
function addProduct() {
  const product = {
    product_name: document.getElementById("name").value,
    manufacturer: document.getElementById("manufacturer").value || "",
    stock: parseInt(document.getElementById("stock").value || "0"),
    price: parseFloat(document.getElementById("price").value || "0"),
    description: document.getElementById("description").value || "",
    category_name: document.getElementById("category").value || ""
  };

  console.log("Adding product:", product);

  fetch("/api/products", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(product)
  })
  .then(res => res.json())
  .then(data => {
    console.log("Server response:", data);
    alert("Product added!");
    renderProduct(product); // Visar produkten direkt på hemsidan 
    
    document.getElementById("name").value = "";
    document.getElementById("manufacturer").value = "";
    document.getElementById("stock").value = "";
    document.getElementById("price").value = "";
    document.getElementById("description").value = "";
    document.getElementById("category").value = "";
  })
  .catch(err => console.error("Error adding product:", err));
}
