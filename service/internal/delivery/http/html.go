package http

import(
	"net/http"
	"fmt"
)

var html = 
`<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Получение заказа по UID</title>
		<script>
			async function fetchOrder() {
				const orderUId = document.getElementById('orderUId').value;
				const response = await fetch('/orders/' + orderUId);
				if (response.ok) {
					const order = await response.json();
					document.getElementById('result').innerText = JSON.stringify(order, null, 2);
				} else {
					document.getElementById('result').innerText = 'Ошибка при получении заказа';
				}
			}
		</script>
	</head>
	<body>
		<h1>Получение заказа по UID</h1>
		<label for="orderId">Введите UID заказа:</label>
		<input type="text" id="orderUId" required>
		<button onclick="fetchOrder()">Получить заказ</button>
		<h2>Результат:</h2>		
		<pre id="result"></pre>	
	</body>
</html>`

func (h *OrderHandler) GetHTMLPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}
