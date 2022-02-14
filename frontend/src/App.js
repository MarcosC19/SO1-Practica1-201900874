import { useState } from 'react';
import './App.css';
import Calculator from './components/calculator';
import TableLogs from './components/tableLogs';

function App() {
  const [reloadOperation, setOperation] = useState(false)
  return (
    <div className="App">
      <Calculator myState={[reloadOperation, setOperation]} />
      <TableLogs myState={[reloadOperation, setOperation]} />
    </div>
  );
}

export default App;
