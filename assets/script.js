document.addEventListener("DOMContentLoaded", function () {
  var preloadElement = document.querySelector(".preload");
  var contentElement = document.querySelector(".content");

  if (preloadElement && contentElement) {
    setTimeout(function () {
      preloadElement.style.opacity = "0";
    }, 2000);

    setTimeout(function () {
      contentElement.style.opacity = "1";
    }, 3000);
  }

  const navLinks = document.querySelectorAll(".navbar a.nav-link");

  navLinks.forEach((link) => {
    link.addEventListener("click", function () {
      navLinks.forEach((link) => link.classList.remove("active"));

      this.classList.add("active");
    });
  });

  const currentHash = window.location.hash;
  const activeLink = document.querySelector(
    `.navbar a.nav-link[href="${currentHash}"]`
  );

  if (activeLink) {
    activeLink.classList.add("active");
  }

  // function to open/close nav
  function toggleNav() {
    // if nav is open, close it
    if (document.querySelector("nav").style.display !== "none") {
      document.querySelector("nav").style.display = "none";
      document.querySelector("button").classList.remove("menu");
    }
    // if nav is closed, open it
    else {
      document.querySelector("button").classList.add("menu");
      document.querySelector("nav").style.display = "flex";
    }
  }

  // when clicking + or ☰ button
  document.querySelector("button").addEventListener("click", function () {
    // when clicking ☰ button, open nav
    if (document.querySelector("header").classList.contains("open")) {
      toggleNav();
    }
    // when clicking + button, open header
    else {
      document.querySelector("header").classList.add("open");
    }
  });

  // close nav
  document.querySelector("#nav-close").addEventListener("click", function () {
    toggleNav();
  });

  // scroll to sections
  document.querySelectorAll("nav li").forEach(function (li) {
    li.addEventListener("click", function () {
      // get index of clicked li and select according section
      var index = Array.from(this.parentNode.children).indexOf(this);
      var target = document.querySelectorAll("content section")[index];

      toggleNav();

      window.scrollTo({
        top: target.offsetTop,
        behavior: "smooth",
      });
    });
  });
});
