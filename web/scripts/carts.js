document.addEventListener("click", function(e) {
  if (e.target.classList.contains("checkout-btn")) {
    fetch("/api/carts/checkout", { 
        method: "POST"
     });
  }
});