import {useEffect, useState} from "react";

const App = () => {
    let expenditures = useState([])
    useEffect(() => {
        fetch("http://localhost:8080/expenditures")
            .then((response) => response.json())
            .then(data => {
                expenditures = data;
                console.log(data);
            })
    }, [])
    return (
        <div className="App">
            <h1>Expenditures</h1>
            <ul>
                {expenditures.map((expenditure) => (
                    <li>
                        {expenditure}
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default App;
