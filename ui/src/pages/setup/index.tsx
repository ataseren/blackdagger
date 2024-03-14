import React from 'react';
import { Grid, Divider, Typography, Button } from '@mui/material';

const Setup = () => {
  const features = [
    {
      buttonText: 'Feature 1',
      paragraphText: 'Description for Feature 1',
    },
    {
      buttonText: 'Feature 2',
      paragraphText: 'Description for Feature 2',
    },
  ];

  return (
    <Grid container spacing={2} style={{ height: '100vh' }}>
      {features.map((feature, index) => (
        <Grid item xs={12} key={index}>
          <Grid container spacing={2} alignItems="center">
            <Grid item xs={3} style={{ paddingLeft: '100px' }}>
              <Button variant="contained" color="primary" style={{padding: '30px 60px', fontSize: '16px',}}>
                {feature.buttonText}
              </Button>
            </Grid>
            <Grid item xs={9}>
              <Typography variant="body1">{feature.paragraphText}</Typography>
            </Grid>
          </Grid>
        </Grid>
      ))}
    </Grid>
  );
};

export default Setup;
