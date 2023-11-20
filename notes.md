const headers = new Headers();
headers.append('Content-Type', 'application/json');
const requestOptions = {
  method: 'GET', 
  headers: headers,
};

fetch('https://go-crud-mongo.onrender.com/employee', requestOptions)
  .then((res) => {
    if (!res.ok) {
      throw new Error('Network response was not ok');
    }
    return res.json();
  })
  .then((data) => {
    console.log(data);
  })
  .catch((err) => {
    console.error('Error:', err);
  });