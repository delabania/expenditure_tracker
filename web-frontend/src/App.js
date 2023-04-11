import {useEffect, useState} from "react";
import {Expenditures} from "./Expenditures";
import Header from "./Header";

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
            <div className={"row"}>
                <Header/>
            </div>
            //todo: fix it
            <div className="space space--lg"/>
            <div className="space space--lg"/>
            <div className={"row"}>
                <div className={"col-6 offset-3"}>
                    <Expenditures/>
                </div>
            </div>
        </div>
    )
}

export default App;
