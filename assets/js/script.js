document.addEventListener("DOMContentLoaded", function () {
  var preloadElement = document.querySelector(".preload");
  var contentElement = document.querySelector(".content");

  // Check if the elements exist on the page
  if (preloadElement && contentElement) {
    // Hide the preload element
    setTimeout(function () {
      preloadElement.style.opacity = "0";
    }, 2000);

    // Fade in the content element after the preload is hidden
    setTimeout(function () {
      contentElement.style.opacity = "1";
    }, 3000); // Adjust the timing as needed
  }

  // Add an event listener to each link to handle the click event
  const navLinks = document.querySelectorAll(".navbar a.nav-link");

  navLinks.forEach((link) => {
    link.addEventListener("click", function () {
      // Remove the 'active' class from all links
      navLinks.forEach((link) => link.classList.remove("active"));

      // Add the 'active' class to the clicked link
      this.classList.add("active");
    });
  });

  // Set the initial active link based on the current hash in the URL
  const currentHash = window.location.hash;
  const activeLink = document.querySelector(
    `.navbar a.nav-link[href="${currentHash}"]`
  );

  if (activeLink) {
    activeLink.classList.add("active");
  }
});
