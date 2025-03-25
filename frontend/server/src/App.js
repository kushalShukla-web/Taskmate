import "./App.css"
import {use, useEffect, useState} from "react"
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid2';
import {TextField, Typography,Checkbox,Box} from "@mui/material";
import DeleteIcon from '@mui/icons-material/Delete';
import AddCircleIcon from '@mui/icons-material/AddCircle';import SendIcon from '@mui/icons-material/Send';
import {blue} from "@mui/material/colors";
import { ThemeProvider, createTheme } from '@mui/material/styles';
import { LocalizationProvider, DatePicker } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";

const theme = createTheme();
function App() {
    const [val ,setval]= useState("")
    const [data,setdata]= useState([])
    const [store ,setstore]= useState([])
    const [desc ,setdescr] = useState("")
    const [date,setdate]= useState("")
    const [show,setshow]= useState(false)
    const handlechange =(e)=>{
        setval(()=>e.target.value)
    }
    const handledesc = (e)=>{
        setdescr(e.target.value)
    }
    const handledate =(e)=>{
      setdate(e)
    }
    // adding data functionality
    const handleClick = async () => {
        try {
            const response = await fetch("http://localhost:8080", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InBpaG9vNTZAZ21haWwuY29tIiwiZXhwIjoxNzQyOTM1MjUwfQ.HNA6D_sKIxnhY_EHKAHbYmaS8CXWBZzh95GoDS9MWJ0"
                },
                body: JSON.stringify(
                    [{
                        "task": `${val}`,
                        "date": `${date}`,
                        "description": `${desc}`,
                        "is_check": true
                    }]
                )
            });
            await getdata()

            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }



        } catch (error) {
            console.error("Error fetching data:", error);
        }
    };
    console.log(store)
    const displaydescription=()=>{
        setdescr(!desc)
    }
    // getting data functionality
    const getdata =()=>{
        fetch("http://localhost:8080",{
            method:"GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InBpaG9vNTZAZ21haWwuY29tIiwiZXhwIjoxNzQyOTM1MjUwfQ.HNA6D_sKIxnhY_EHKAHbYmaS8CXWBZzh95GoDS9MWJ0"
            }
        }).then(data=>{ return data.json()})
            .then((data)=>setstore(data))
            .catch(error => console.error("Error fetching data:", error));
    }
    useEffect(() => {
        getdata()
    },[]);

    // delete functionality
 const handclickdelete=  async (index)=>{
     console.log(index)
     fetch("http://localhost:8080",{
         method: "DELETE",
         headers:{
             "Content-Type": "application/json",
             "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InBpaG9vNTZAZ21haWwuY29tIiwiZXhwIjoxNzQyOTM1MjUwfQ.HNA6D_sKIxnhY_EHKAHbYmaS8CXWBZzh95GoDS9MWJ0"
         },
         body: JSON.stringify(
             [{
                 id: index
             }]
         )
     })
     console.log("Data deleted")
    await  getdata()

 }
 const handleshow=()=>{
     setshow(!show)
 }

  return (
      <>
          <Typography variant="h3" align="center" sx={{ mt: 2 }}>
              Todo App
          </Typography>
          <Box sx={{
              display:"flex",
              alignItems: "center",
              justifyContent:"center",
              height:"100vh",
          }}>
              <Typography>
                  {
                      <Box sx={{ mt: 3 }}>  {/* mb = margin-bottom (3px) */}
                          {store.map((val, index) => (
                              <Grid container spacing={2} sx={{ padding: "10px" }} key={index}>
                                  <Grid item xs={10}>
                                      <TextField value={val.task} />
                                  </Grid>
                                  <Grid item xs={2}>
                                      <DeleteIcon onClick={() => {handclickdelete(val.id)}}/>
                                  </Grid>
                              </Grid>
                          ))}
                      </Box>
                  }
                  {
                      show?<Grid container spacing={2} >
                          <Grid>
                              <TextField
                              label='taskk'
                              onChange={(event)=>{handlechange(event)}}
                              />
                          </Grid>
                        <Grid xs={12}>
                         <TextField
                         label='description'
                         value={desc}
                         onChange={(event)=>{handledesc(event)}}
                         />
                          </Grid>
                          <Grid xs={12}>
                              <LocalizationProvider dateAdapter={AdapterDayjs}>
                                  <DatePicker
                                      label="Select Date"
                                      renderInput={(params) => <TextField {...params} />}
                                      onChange={handledate}
                                  />
                              </LocalizationProvider>
                          </Grid>
                      </Grid>:null
                  }
                  {show?<SendIcon
                      onClick={()=>{handleClick()}}
                      />:<></>
                  }
                  {show?<></>:<AddCircleIcon
                  sx={{
                      paddingLeft:"160px",
                      fontSize: "80px",
                      color: "skyblue"
                  }}
                  onClick={()=>{handleshow()}}
                  />}

              </Typography>
          </Box>
      </>

  );
}

export default App;
