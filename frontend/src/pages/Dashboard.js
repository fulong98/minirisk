import React from 'react';
import { Grid, Paper, Typography, Box } from '@mui/material';
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from 'recharts';
import { useMargin } from '../contexts/MarginContext';

function Dashboard() {
  const { marginData } = useMargin();

  // Sample data for demonstration
  const portfolioData = [
    { name: 'AAPL', value: 150000 },
    { name: 'GOOGL', value: 200000 },
    { name: 'MSFT', value: 180000 },
    { name: 'AMZN', value: 220000 },
  ];

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Dashboard
      </Typography>
      <Grid container spacing={3}>
        {/* Summary Cards */}
        <Grid item xs={12} sm={6} md={3}>
          <Paper sx={{ p: 2 }}>
            <Typography color="textSecondary" gutterBottom>
              Total Portfolio Value
            </Typography>
            <Typography variant="h5">
              ${marginData?.portfolioValue?.toLocaleString() || '0'}
            </Typography>
          </Paper>
        </Grid>
        <Grid item xs={12} sm={6} md={3}>
          <Paper sx={{ p: 2 }}>
            <Typography color="textSecondary" gutterBottom>
              Net Equity
            </Typography>
            <Typography variant="h5">
              ${marginData?.netEquity?.toLocaleString() || '0'}
            </Typography>
          </Paper>
        </Grid>
        <Grid item xs={12} sm={6} md={3}>
          <Paper sx={{ p: 2 }}>
            <Typography color="textSecondary" gutterBottom>
              Margin Shortfall
            </Typography>
            <Typography
              variant="h5"
              color={marginData?.marginShortfall > 0 ? 'error' : 'success'}
            >
              ${marginData?.marginShortfall?.toLocaleString() || '0'}
            </Typography>
          </Paper>
        </Grid>
        <Grid item xs={12} sm={6} md={3}>
          <Paper sx={{ p: 2 }}>
            <Typography color="textSecondary" gutterBottom>
              Margin Call Status
            </Typography>
            <Typography
              variant="h5"
              color={marginData?.marginCall ? 'error' : 'success'}
            >
              {marginData?.marginCall ? 'Active' : 'Clear'}
            </Typography>
          </Paper>
        </Grid>

        {/* Charts */}
        <Grid item xs={12} md={8}>
          <Paper sx={{ p: 2, height: 400 }}>
            <Typography variant="h6" gutterBottom>
              Portfolio Distribution
            </Typography>
            <ResponsiveContainer width="100%" height="100%">
              <BarChart data={portfolioData}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="name" />
                <YAxis />
                <Tooltip />
                <Legend />
                <Bar dataKey="value" fill="#1976d2" />
              </BarChart>
            </ResponsiveContainer>
          </Paper>
        </Grid>
        <Grid item xs={12} md={4}>
          <Paper sx={{ p: 2, height: 400 }}>
            <Typography variant="h6" gutterBottom>
              Risk Metrics
            </Typography>
            {/* Add risk metrics visualization here */}
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
}

export default Dashboard; 