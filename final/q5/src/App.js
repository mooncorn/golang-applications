import { useState } from 'react';
import PatientForm from './components/PatientForm';
import PatientTable from './components/PatientTable'

function App() {
  const [patients, setPatients] = useState({})
  const [searchResults, setSearchResults] = useState({})

  return (
    <>
      <PatientForm patients={patients} setPatients={setPatients} setSearchResults={setSearchResults}/>

      <h1 className="mt-2">Search Results</h1>
      <PatientTable patients={searchResults}/>

      <h1 className="mt-2">All Patients</h1>
      <PatientTable patients={patients}/>
    </>
  );
}

export default App;
