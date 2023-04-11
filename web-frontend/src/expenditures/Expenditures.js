import React from "react";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faShoppingBasket} from "@fortawesome/free-solid-svg-icons";


class Expenditures extends React.Component {
    render() {
        return (
            <div className={"row"}>
                <h6>Expenditures</h6>
                <table className={"table striped"}>
                    <thead>
                    <tr className={"uppercase"}>
                        <th>Date</th>
                        <th>Description</th>
                        <th>Category</th>
                        <th>Owner</th>
                        <th>Amount</th>
                        <th></th>
                    </tr>
                    </thead>
                    <tbody>
                    <Expenditure/>
                    <Expenditure/>
                    <Expenditure/>
                    <Expenditure/>
                    <Expenditure/>
                    <Expenditure/>
                    </tbody>
                </table>
            </div>
        )
    }
}

class Expenditure extends React.Component {
    render() {
        return (
            <tr className={"tex-centered"}>
                <td>
                    2022-01-01
                </td>
                <td>
                    Żabka - zakupy
                </td>
                <td>
                    <FontAwesomeIcon icon={faShoppingBasket} color={"green"}/>
                </td>
                <td>
                    Przemo
                </td>
                <td>
                    100.00
                </td>
                <td>
                </td>
            </tr>
        )
    }
}


export {Expenditure, Expenditures}