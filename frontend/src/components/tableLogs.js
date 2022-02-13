import React from "react";
import '../css/tableLogs.css'
import { Table } from 'reactstrap'

function TableLogs() {
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
                </tbody>
            </Table>
        </div>
    )
}

export default TableLogs