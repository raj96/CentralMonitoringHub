import { Add } from "@mui/icons-material";
import {
  Button,
  Divider,
  Grid,
  MenuItem,
  Paper,
  Select,
  Table,
  TableContainer,
  TableHead,
  TableCell,
  TextField,
  Typography,
  TableRow,
  TableBody,
  InputLabel,
  FormControl,
} from "@mui/material";
import { Box } from "@mui/system";
import axios from "axios";
import { useEffect, useRef, useState } from "react";
import { v4 as uuid } from "uuid";

export default function SourcePage() {
  const name = useRef("");
  const allowedMachines = useRef("");
  const srcTypeName = useRef("");

  const [sourceData, setSourceData] = useState([]);
  const [srcTypeNameList, setSrcTypeNameList] = useState([]);
  const [hideError, setHideError] = useState(true);
  const [errorMsg, setErrorMsg] = useState("");

  const showErrorMsg = (err) => {
    if (err.response) {
      console.error(err.response.data);
      setErrorMsg(err.response.data.error);
      setHideError(false);
    }
  };

  const fetchList = () => {
    axios
      .get("/fetchlist?id=test&listName=source")
      .then((res) => {
        setSourceData(res?.data);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const fetchSourceTypeNames = () => {
    axios
      .get("/fetchlist?id=test&listName=sourcetype")
      .then((res) => {
        setSrcTypeNameList(res?.data?.map((data) => data.name));
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const addSource = () => {
    console.log(name.current)
    const nameValue = name.current.value;
    const allowedMachinesValue = allowedMachines.current.value;
    const srcTypeNameValue = srcTypeName.current.value;
    const postData = {
      id: "test",
      name: nameValue,
      allowedMachines: allowedMachinesValue,
      sourceTypeName: srcTypeNameValue,
    };
    console.log(postData)

    axios
      .post("/addsource", postData)
      .then((res) => {
        if (res.status === 200) {
          fetchList();
        }
      })
      .catch((err) => {
        showErrorMsg(err);
      });
  };

  useEffect(() => {
    fetchSourceTypeNames();
    fetchList();
  }, []);

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
            inputRef={allowedMachines}
            label="Allowed Machines(Optional)"
            size="small"
            fullWidth
          />
        </Grid>
        <Grid item xs={6}>
          <FormControl fullWidth>
            <InputLabel id="src-type-name-label">Source Type name *</InputLabel>
            <Select
              inputRef={srcTypeName}
              labelId="src-type-name-label"
              label="Source Type name *"
              size="small"
              fullWidth
              required
            >
              {srcTypeNameList?.map((name) => (
                <MenuItem key={uuid()} value={`${name}`}>
                  {name}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>
        <Grid item xs={6}>
          <Button
            fullWidth
            variant="contained"
            startIcon={<Add />}
            onClick={addSource}
          >
            <Typography variant="body1">Source</Typography>
          </Button>
        </Grid>
        <Grid item xs={12}>
          <Typography variant="caption" hidden={hideError} color="red">
            {errorMsg}
          </Typography>
        </Grid>
      </Grid>
      <Divider />
      <TableContainer component={Paper} sx={{ marginTop: 2 }}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell align="center">Name</TableCell>
              <TableCell align="center">Endpoint</TableCell>
              <TableCell align="center">Source Type</TableCell>
              <TableCell align="center">Allowed Machines</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {sourceData?.map((data) => (
              <TableRow key={uuid()}>
                <TableCell align="center">{data.name}</TableCell>
                <TableCell align="center">/{data.id}</TableCell>
                <TableCell align="center">{data.sourcetypename}</TableCell>
                <TableCell align="center">{data.allowedmachines}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
}
