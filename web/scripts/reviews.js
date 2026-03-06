function loadReviews(productId) {

  fetch("/api/reviews/" + productId)
    .then(res => res.json())
    .then(reviews => {

      const container = document.getElementById("reviews-" + productId);

      container.innerHTML = "";   // viktigt

      reviews.forEach(r => {

        const div = document.createElement("div");

        div.innerHTML = `
          <p>⭐ ${r.grade}/5</p>
          <p>${r.comment_text}</p>
        `;

        container.appendChild(div);

      });

    });
}