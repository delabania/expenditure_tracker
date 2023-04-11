import {useEffect, useState} from "react";
import Expenditure from "./expenditures";

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
            <table>
                <tr>
                    <th>Date</th>
                    <th>Category</th>
                    <th>Amount</th>
                    <th>Description</th>
                    <th>Owner</th>
                </tr>
                {expenditures.map((expenditure) => (
                    <Expenditure expenditure={expenditure}/>
                ))}
            </table>
        </div>
    )
}

export default App;
