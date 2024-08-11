document.addEventListener('DOMContentLoaded', function() {
  console.log("JavaScript is working!");
  const heading = document.querySelector('h1');
  if (heading) {
      heading.style.color = 'red';
  }
});
