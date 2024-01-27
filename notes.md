const headers = new Headers();
headers.append('Content-Type', 'application/json');
const requestOptions = {
method: 'GET',
headers: headers,
};

const example = fetch('https://go-crud-todo.onrender.com/getAllTodos', requestOptions)
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

console.log({example})
