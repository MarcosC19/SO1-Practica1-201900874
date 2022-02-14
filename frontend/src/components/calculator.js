import React, { useState } from "react";
import '../css/calculator.css'

function Calculator() {
    const [values, setValues] = useState({
        num1: '',
        num2: '',
        operation: '',
        result: ''
    })

    function changeValue(e) { // USE TO SET IN NUMBER1 THE DATA STRING
        const lastValue = values
        var newValue = values.num2 + e.target.value
        setValues({
            num1: lastValue.num1,
            num2: newValue,
            operation: lastValue.operation,
            result: ''
        })
    }

    function setOperation(e) { // USE TO SET AN OPERATION IN THE STATE
        const lastValue = values
        setValues({
            num1: lastValue.num2,
            num2: '',
            operation: e.target.value,
            result: ''
        })
    }

    function getResult() { // USE TO GET THE RESULT WITH THE BACKEND
        const lastJSON = {
            num1: parseFloat(values.num1),
            num2: parseFloat(values.num2),
            operation: values.operation
        }
        fetch('http://localhost:5000/Operation', {
            method: 'POST',
            body: JSON.stringify(lastJSON)
        })
            .then(res => res.json())
            .then(data => {
                setValues({
                    num1: '',
                    num2: data.result.toString(),
                    operation: '',
                    result: ''
                })
            })
    }

    function restartValues() { // USE TO SET DEFAULT VALUES STATE
        setValues({
            num1: '',
            num2: '',
            operation: '',
            result: ''
        })
    }

    return (
        <div className="contentCalculator">
            <input type="text" id="texto" value={values.num2} disabled />
            <button className="fila0" id='cero' value='0' onClick={changeValue}>0</button>
            <button className="fila0" id='punto' value='.' onClick={changeValue}>.</button>
            <button className="fila0" id='igual' onClick={getResult}>=</button>
            <button className="fila1" id='col1' value='1' onClick={changeValue}>1</button>
            <button className="fila1" id='col2' value='2' onClick={changeValue}>2</button>
            <button className="fila1" id='col3' value='3' onClick={changeValue}>3</button>
            <button className="fila2" id='col1' value='4' onClick={changeValue}>4</button>
            <button className="fila2" id='col2' value='5' onClick={changeValue}>5</button>
            <button className="fila2" id='col3' value='6' onClick={changeValue}>6</button>
            <button className="fila3" id='col1' value='7' onClick={changeValue}>7</button>
            <button className="fila3" id='col2' value='8' onClick={changeValue}>8</button>
            <button className="fila3" id='col3' value='9' onClick={changeValue}>9</button>
            <button className="fila0" id='division' value='/' onClick={setOperation}>/</button>
            <button className="fila4" id="col14" value='+' onClick={setOperation}>+</button>
            <button className="fila4" id="col24" value='-' onClick={setOperation}>-</button>
            <button className="fila4" id="col34" value='*' onClick={setOperation}>*</button>
            <button className="fila4" id="col4" onClick={restartValues}>DEL</button>
        </div>
    )
}

export default Calculator