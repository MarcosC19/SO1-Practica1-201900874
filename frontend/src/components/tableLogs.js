import React, { useState, useEffect } from "react";
import '../css/tableLogs.css'
import { Table } from 'reactstrap'

function TableLogs(props) {
    const [operations, setOperations] = useState({
        operations: []
    })

    const [updateOpt, setUpdateOpt] = props.myState

    function getOperations() {
        fetch('http://localhost:5000/getOperations')
            .then(res => res.json())
            .then(data => {
                if (data.data != null) {
                    setOperations({
                        operations: data.data
                    })
                }
            })
    }

    useEffect(() => {
        getOperations()
        setUpdateOpt(false)
    }, [updateOpt, setUpdateOpt])

    return (
        <div className="contentTable">
            <Table>
                <thead>
                    <tr>
                        <th>Numero1</th>
                        <th>Numero2</th>
                        <th>Operacion</th>
                        <th>Resultado</th>
                        <th>Fecha y hora</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        operations.operations.map(operation => {
                            var dateFinal = new Date(operation.Date)
                            return (
                                <tr key={operation.ID}>
                                    <td>{operation.Number1}</td>
                                    <td>{operation.Number2}</td>
                                    <td>{operation.Operation}</td>
                                    <td>{operation.Result}</td>
                                    <td>{dateFinal.toLocaleString()}</td>
                                </tr>
                            )
                        })
                    }
                </tbody>
            </Table>
        </div>
    )
}

export default TableLogs