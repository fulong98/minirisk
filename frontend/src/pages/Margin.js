import React from 'react';
import {
  Box,
  Typography,
  Paper,
  Grid,
  Alert,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@mui/material';
import { useMargin } from '../contexts/MarginContext';

function Margin() {
  const { marginData, loading, error } = useMargin();

  if (loading) {
    return (
      <Box>
        <Typography variant="h4" gutterBottom>
          Margin Status
        </Typography>
        <Typography>Loading...</Typography>
      </Box>
    );
  }

  if (error) {
    return (
      <Box>
        <Typography variant="h4" gutterBottom>
          Margin Status
        </Typography>
        <Alert severity="error">Error loading margin data: {error}</Alert>
      </Box>
    );
  }

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Margin Status
      </Typography>

      {marginData?.marginCall && (
        <Alert severity="error" sx={{ mb: 2 }}>
          MARGIN CALL ALERT: Please deposit additional funds to meet margin requirements
        </Alert>
      )}

      <Grid container spacing={3}>
        <Grid item xs={12} md={6}>
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Margin Requirements
            </Typography>
            <TableContainer>
              <Table>
                <TableBody>
                  <TableRow>
                    <TableCell>Portfolio Value</TableCell>
                    <TableCell align="right">
                      ${marginData?.portfolioValue?.toLocaleString()}
                    </TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell>Net Equity</TableCell>
                    <TableCell align="right">
                      ${marginData?.netEquity?.toLocaleString()}
                    </TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell>Margin Shortfall</TableCell>
                    <TableCell
                      align="right"
                      sx={{
                        color: marginData?.marginShortfall > 0 ? 'error.main' : 'success.main',
                      }}
                    >
                      ${marginData?.marginShortfall?.toLocaleString()}
                    </TableCell>
                  </TableRow>
                </TableBody>
              </Table>
            </TableContainer>
          </Paper>
        </Grid>

        <Grid item xs={12} md={6}>
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Risk Metrics
            </Typography>
            <TableContainer>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Metric</TableCell>
                    <TableCell align="right">Value</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  <TableRow>
                    <TableCell>Initial Margin Requirement</TableCell>
                    <TableCell align="right">25%</TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell>Maintenance Margin Requirement</TableCell>
                    <TableCell align="right">20%</TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell>Current Margin Ratio</TableCell>
                    <TableCell align="right">
                      {(
                        (marginData?.netEquity / marginData?.portfolioValue) *
                        100
                      ).toFixed(2)}
                      %
                    </TableCell>
                  </TableRow>
                </TableBody>
              </Table>
            </TableContainer>
          </Paper>
        </Grid>

        <Grid item xs={12}>
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Margin Call History
            </Typography>
            <TableContainer>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Date</TableCell>
                    <TableCell>Amount</TableCell>
                    <TableCell>Status</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  <TableRow>
                    <TableCell>No margin calls recorded</TableCell>
                    <TableCell>-</TableCell>
                    <TableCell>-</TableCell>
                  </TableRow>
                </TableBody>
              </Table>
            </TableContainer>
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
}

export default Margin; 