import {useEffect, useState} from "react";
import {Expenditures} from "./expenditures/Expenditures";
import Header from "./Header";
import Summary from "./expenditures/Summary";

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
                <div className={"col-6 offset-2"}>
                    <Expenditures/>
                    <div className={"row"}>
                        <div className={"col-3 offset-9"}>
                            <button className={"btn-layout btn-info btn--md btn--pilled py-1 px-2"}>Add expenditure</button>
                        </div>
                    </div>
                </div>
            </div>

            <div className={"row"}>
                <div className={"col-6 offset-2"}>
                    <Summary/>
                </div>
            </div>

        </div>
    )
}

export default App;
