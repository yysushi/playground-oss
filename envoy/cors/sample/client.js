fetch("http://localhost:3000/hello")
  .then((response) => response.json())
  .then((data) => console.log(data));
