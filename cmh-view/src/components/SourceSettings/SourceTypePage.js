import { Add } from "@mui/icons-material";
import {
  Button,
  Divider,
  Grid,
  Paper,
  Table,
  TableContainer,
  TableHead,
  TableCell,
  TextField,
  Typography,
  TableRow,
  TableBody,
} from "@mui/material";
import { Box } from "@mui/system";
import axios from "axios";
import { useEffect, useRef, useState } from "react";
import {v4 as uuid} from "uuid";

export default function SourceTypePage() {
  const [sourceTypeData, setSourceTypeData] = useState([]);
  const name = useRef();
  const format = useRef();

  const loadSourceTypes = () => {
    axios
      .get("/fetchlist?id=test&listName=sourcetype")
      .then((res) => {
        setSourceTypeData(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const addSourceType = () => {
    const stats = {};
    for(let fields of  format.current.value.split(",")) {
      const group = fields.split(":")
      stats[group[0]] = group[1]
    }
    const postData = {
      name: name.current.value,
      stats
    };

    axios
    .post("/addsourcetype", postData)
    .then(loadSourceTypes)
    .catch(alert);
  };


  useEffect(loadSourceTypes, []);
  
  return (
    <Box sx={{ p: 2 }}>
      <Grid
        container
        spacing={2}
        justifyContent="space-evenly"
        alignItems="center"
        sx={{ marginBottom: 2 }}
      >
        <Grid item xs={6}>
          <TextField inputRef={name} label="Name" required size="small" fullWidth />
        </Grid>
        <Grid item xs={6}>
          <TextField
            inputRef={format}
            label="Stats structure"
            placeholder="format: cpu:number,dmesg:log,..."
            required
            size="small"
            fullWidth
          />
        </Grid>
        <Grid item xs={6}>
          <Button fullWidth variant="contained" startIcon={<Add />} onClick={addSourceType}>
            <Typography variant="body1">Sourcetype</Typography>
          </Button>
        </Grid>
      </Grid>
      <Divider />
      <TableContainer component={Paper} sx={{ marginTop: 2 }}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell align="center">Name</TableCell>
              <TableCell align="center">Stats Structure</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {sourceTypeData?.map((data) => {
              var statsString = "";
              for (var key in data.stats) {
                statsString += `${key}:${data.stats[key]},`;
              }
              statsString = statsString.slice(0, -1);
              return (
                <TableRow key={uuid()}>
                  <TableCell align="center">{data.name}</TableCell>
                  <TableCell align="center">{statsString}</TableCell>
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
}
