let request = () => {
    result.innerHTML = '';
    axios.get('/api/scrap?url=' + input.value)
        .then(resp => {
            let table = document.createElement('table');
            table.style = 'width: 500px';
            let tbody = document.createElement('tbody');

            resp.data.forEach(x => {
                let tr = document.createElement('tr');
                let name = document.createElement('td');
                name.style = "vertical-align: top; font-style: italic";
                name.innerText = x.name + ':';
                let value = document.createElement('td');
                value.innerText = x.value;
                tr.appendChild(name);
                tr.appendChild(value);
                tbody.appendChild(tr);
            })

            table.appendChild(tbody);

            let b = document.createElement('b');
            b.innerText = "Website info"

            result.appendChild(b);
            result.appendChild(table);
        })
        .catch(x => {
            let b = document.createElement('b');
            b.innerText = "Error occured";

            let status = document.createElement('span');
            status.innerText = "Code - " + x.response.status;

            let message = document.createElement('span');
            message.innerText = "Message - " + x.response.data;

            result.appendChild(b);
            result.appendChild(status);
            result.appendChild(message);
        });
}
