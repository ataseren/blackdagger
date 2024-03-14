import React, { useState , PureComponent } from 'react';
import { Container, Grid, Typography, TextField, Button } from '@mui/material';
import ReactDiffViewer from 'react-diff-viewer';

function DiffPage() {
  const [oldFileContent, setOldFileContent] = useState<string>('');
  const [newFileContent, setNewFileContent] = useState<string>('');

  const handleOldFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      const reader = new FileReader();
      reader.onload = () => {
        const content = reader.result as string;
        setOldFileContent(content);
      };
      reader.readAsText(event.target.files[0]);
    }
  };

  const handleNewFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      const reader = new FileReader();
      reader.onload = () => {
        const content = reader.result as string;
        setNewFileContent(content);
      };
      reader.readAsText(event.target.files[0]);
    }
  };

  return (
    <Container maxWidth="xl">
      <Grid container spacing={2} justifyContent="center" alignItems="flex-start" style={{ flexGrow: 1 }}>
        <Grid item xs={12}>
          <Typography variant="h4" align="center">
            Upload Files
          </Typography>
        </Grid>
        <Grid item xs={6}>
          <TextField
            type="file"
            variant="outlined"
            fullWidth
            onChange={handleOldFileChange}
          />
        </Grid>
        <Grid item xs={6}>
          <TextField
            type="file"
            variant="outlined"
            fullWidth
            onChange={handleNewFileChange}
          />
        </Grid>
        {oldFileContent && newFileContent && (
          <Grid item xs={12}>
            <Typography variant="h4" align="center">
              File Differences
            </Typography>
            <div>
              <ReactDiffViewer
                oldValue={oldFileContent}
                newValue={newFileContent}
                splitView={true}
                useDarkTheme={true}
              />
            </div>
          </Grid>
        )}
      </Grid>
    </Container>
  );
}

export default DiffPage;
