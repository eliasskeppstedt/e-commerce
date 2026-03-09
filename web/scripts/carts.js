document.addEventListener("click", function (e) {
  if (e.target.classList.contains("checkout-btn")) {
    fetch("/api/carts/checkout", {
      method: "POST"
    });
  }
  else if (e.target.classList.contains("remove-btn")) {
    fetch("/api/carts/items", {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        product_id: 1//Number(productId)
      })
    });
  }
});