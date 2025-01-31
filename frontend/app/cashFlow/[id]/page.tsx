import React from "react";
export default function transaction({params} : {params : {id : string}}){
    const {id}= params
    return(
        <p>
            This is the transactions of that specific cashflow
        </p>
    )
}