import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Paper,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
} from '@mui/material';
import { DataGrid } from '@mui/x-data-grid';
import axios from 'axios';

function Positions() {
  const [positions, setPositions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [open, setOpen] = useState(false);
  const [newPosition, setNewPosition] = useState({
    symbol: '',
    quantity: 0,
    costBasis: 0,
  });

  const columns = [
    { field: 'id', headerName: 'ID', width: 90 },
    { field: 'symbol', headerName: 'Symbol', width: 130 },
    { field: 'quantity', headerName: 'Quantity', width: 130 },
    {
      field: 'costBasis',
      headerName: 'Cost Basis',
      width: 130,
      valueFormatter: (params) => `$${params.value.toFixed(2)}`,
    },
    {
      field: 'marketValue',
      headerName: 'Market Value',
      width: 130,
      valueFormatter: (params) => `$${params.value.toFixed(2)}`,
    },
    {
      field: 'pnl',
      headerName: 'P&L',
      width: 130,
      valueFormatter: (params) => `$${params.value.toFixed(2)}`,
      cellClassName: (params) =>
        params.value >= 0 ? 'positive-pnl' : 'negative-pnl',
    },
  ];

  const fetchPositions = async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/positions/1');
      setPositions(response.data);
    } catch (error) {
      console.error('Error fetching positions:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPositions();
  }, []);

  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
    setNewPosition({
      symbol: '',
      quantity: 0,
      costBasis: 0,
    });
  };

  const handleSubmit = async () => {
    try {
      await axios.post('http://localhost:8080/api/positions', {
        ...newPosition,
        clientId: 1,
      });
      fetchPositions();
      handleClose();
    } catch (error) {
      console.error('Error creating position:', error);
    }
  };

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Positions
      </Typography>
      <Box sx={{ mb: 2 }}>
        <Button variant="contained" color="primary" onClick={handleOpen}>
          Add Position
        </Button>
      </Box>
      <Paper sx={{ height: 400, width: '100%' }}>
        <DataGrid
          rows={positions}
          columns={columns}
          pageSize={5}
          rowsPerPageOptions={[5]}
          loading={loading}
        />
      </Paper>

      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>Add New Position</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            margin="dense"
            label="Symbol"
            fullWidth
            value={newPosition.symbol}
            onChange={(e) =>
              setNewPosition({ ...newPosition, symbol: e.target.value })
            }
          />
          <TextField
            margin="dense"
            label="Quantity"
            type="number"
            fullWidth
            value={newPosition.quantity}
            onChange={(e) =>
              setNewPosition({ ...newPosition, quantity: parseInt(e.target.value) })
            }
          />
          <TextField
            margin="dense"
            label="Cost Basis"
            type="number"
            fullWidth
            value={newPosition.costBasis}
            onChange={(e) =>
              setNewPosition({
                ...newPosition,
                costBasis: parseFloat(e.target.value),
              })
            }
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleSubmit} color="primary">
            Add
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}

export default Positions; 