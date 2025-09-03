# Code Coverage Setup

This project uses Codecov for code coverage reporting. This document explains how to set up coverage reporting for your repository.

## Setting up Codecov

1. **Sign up for Codecov**
   - Go to https://codecov.io
   - Sign up using your GitHub account
   - Codecov is free for open source projects

2. **Add your repository**
   - After signing up, Codecov will ask you to select repositories to add
   - Select your `learn-go-api` repository
   - Codecov will automatically configure the repository

3. **GitHub Actions Integration**
   - The project already includes a GitHub Actions workflow (`.github/workflows/test-coverage.yml`) that:
     - Runs tests with coverage
     - Uploads coverage reports to Codecov
   - This workflow runs on every push to the `main` branch and on pull requests

4. **Codecov Token (Optional for Public Repositories)**
   - For public repositories, no token is required
   - For private repositories, you'll need to:
     - Get your repository upload token from Codecov
     - Add it as a secret in your GitHub repository settings:
       1. Go to your repository settings on GitHub
       2. Navigate to "Settings" → "Secrets and variables" → "Actions"
       3. Click "New repository secret"
       4. Name it `CODECOV_TOKEN` and paste the token value
       5. Update the workflow file to use the token:
          ```yaml
          - name: Upload coverage to Codecov
            uses: codecov/codecov-action@v4
            with:
              file: ./coverage.txt
              flags: unittests
              name: codecov-learn-api
              token: ${{ secrets.CODECOV_TOKEN }}
          ```

## Viewing Coverage Reports

Once set up, coverage reports will be available at:
- `https://codecov.io/gh/your-username/learn-go-api`

You can also add a coverage badge to your README:
```markdown
[![codecov](https://codecov.io/gh/your-username/learn-go-api/branch/main/graph/badge.svg)](https://codecov.io/gh/your-username/learn-go-api)
```

## Local Coverage Testing

You can also generate coverage reports locally:

```bash
# Run tests with coverage
go test -coverprofile=coverage.txt ./tests/handlers ./tests/services

# View coverage in terminal
go tool cover -func=coverage.txt

# Generate HTML coverage report
go tool cover -html=coverage.txt -o coverage.html
```

## Troubleshooting

If you encounter issues:

1. **Check the GitHub Actions logs** - Look at the workflow run logs to see any errors
2. **Verify the workflow file** - Ensure the paths and commands are correct
3. **Check Codecov settings** - Make sure your repository is properly configured in Codecov
4. **Repository permissions** - Ensure GitHub Actions has the necessary permissions