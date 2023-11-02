import { useState, useEffect } from "react";
import "./map.css";

const Map = () => {

  const [map, setMap] = useState([{}])
  // const [rows, setRows] = useState([])
  const mapAPI = "http://localhost:3000/map"
  useEffect(()=>{
    fetch(mapAPI)
      .then(res=>res.json())
      .then(data=>{
        setMap(data)
      })
      .catch(err=>{
        console.log(err)
      })
  },[])

  useEffect(()=>{
    console.log(map)
    
  },[map])

  const [hoveredCell, setHoveredCell] = useState(null);

  const handleMouseEnter = (cellId) => {
    setHoveredCell(cellId);
  };

  const handleMouseLeave = () => {
    setHoveredCell(null);
  };

  const createCellId = (row, col) => `cell-${row}-${col}`;

  
  const genRows = () => {
    let rows = [];
    let count = 0
    for (let i = 0; i < 10; i++) {
      const cells = [];
      for (let j = 0; j < 10; j++) {
        
        const cellId = createCellId(i, j);
        const cellStyle = (i + j) % 2 === 0 ? "white" : "lightgrey";
        const isHovered = cellId === hoveredCell;

        let buildingName = "";
        if (map[count] && map[count].BuildingName) {
          buildingName = map[count].BuildingName;
        }
        count += 1
        cells.push(
          <div
            key={cellId}
            className={`cell ${isHovered ? "border-black" : ""}`}
            style={{ background: cellStyle }}
            onMouseEnter={() => handleMouseEnter(cellId)}
            onMouseLeave={handleMouseLeave}
          >{buildingName}</div>
        );
      }
      rows.push(<div key={i} className="row">{cells}</div>);
    }
    return rows
  }

  

  return (
    <div className="etch-a-sketch-grid">
      {genRows()}
    </div>
  );
};

export default Map;