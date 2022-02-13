import React from "react";
import '../css/calculator.css'

function Calculator() {
    return (
        <div className="contentCalculator">
            <input type="text" id="texto" disabled />
            <button className="fila0" id='cero'>0</button>
            <button className="fila0" id='punto'>.</button>
            <button className="fila0" id='igual'>=</button>
            <button className="fila1" id='col1'>1</button>
            <button className="fila1" id='col2'>2</button>
            <button className="fila1" id='col3'>3</button>
            <button className="fila2" id='col1'>4</button>
            <button className="fila2" id='col2'>5</button>
            <button className="fila2" id='col3'>6</button>
            <button className="fila3" id='col1'>7</button>
            <button className="fila3" id='col2'>8</button>
            <button className="fila3" id='col3'>9</button>
            <button className="fila0" id='division'>/</button>
            <button className="fila4" id="col14">+</button>
            <button className="fila4" id="col24">-</button>
            <button className="fila4" id="col34">*</button>
            <button className="fila4" id="col4">DEL</button>
        </div>
    )
}

export default Calculator