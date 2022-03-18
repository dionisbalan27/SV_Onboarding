import React, {useEffect, useState} from 'react';
import "./App.css"
import {Table} from "antd";
import axios from 'axios';


function App() {
  const [allData, setAllData] = useState([]);

  useEffect(() => {
    axios.get(`http://localhost:8001/products`).then(res => {
      setAllData(res.data);
    });
  }, []);

  const columns = [
 
    {
    
      title: 'Name',
      dataIndex: 'name',
    },
    {

      title: 'Description',
      dataIndex: 'description'
    },
    {
    
      title: 'Status',
      dataIndex: 'status'
    },
  ];

  const data = [{}];

  allData.map((user) => {
    data.push({
     key: user.id,
     name: user.Name,
     description: user.Description,
     status: user.Status,
    })
   return data;
 });

 console.log(data)

  return (
    <div className="App">
      <header className="App-Header">
<Table columns={columns} dataSource={data}>

</Table>
console.log
      </header>
    </div>
  );
}

export default App;
