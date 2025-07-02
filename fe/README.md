# 📊 Another Jira Frontend

A modern React-based frontend application for the Another Jira project management system.

## 🏗️ Project Structure

```
fe/
├── public/              # Static files
├── src/                # Source code
│   ├── components/     # Reusable components
│   ├── pages/         # Page components
│   ├── hooks/         # Custom React hooks
│   ├── services/      # API services
│   ├── store/         # Redux store
│   ├── utils/         # Utility functions
│   ├── styles/        # Global styles
│   ├── App.js         # Root component
│   └── index.js       # Entry point
├── package.json       # Dependencies
└── Dockerfile         # Docker configuration
```

## 🚀 Features

- **Modern UI**: Clean and intuitive user interface
- **Responsive Design**: Works on all devices
- **Real-time Updates**: Live task updates
- **Drag and Drop**: Intuitive task management
- **Advanced Filtering**: Powerful search and filter capabilities
- **Dark Mode**: Support for light and dark themes
- **Offline Support**: Progressive Web App features

## ⚙️ Tech Stack

- **Framework**: React 18
- **State Management**: Redux Toolkit
- **UI Library**: Material-UI
- **Routing**: React Router
- **HTTP Client**: Axios
- **Form Handling**: React Hook Form
- **Testing**: Jest + React Testing Library
- **Build Tool**: Create React App
- **Container**: Docker

## 🛠 Setup and Installation

### Prerequisites

- Node.js 16+
- npm or yarn
- Docker (optional)

### Local Development

1. **Clone and Setup**
   ```bash
   git clone https://github.com/your-org/another-jira.git
   cd another-jira/fe
   npm install
   ```

2. **Environment Configuration**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Start Development Server**
   ```bash
   npm start
   ```

### Docker Deployment

```bash
# Build the image
docker build -t another-jira-frontend .

# Run the container
docker run -p 3000:3000 another-jira-frontend
```

## 📚 Available Scripts

- `npm start` - Runs the app in development mode
- `npm test` - Launches the test runner
- `npm run build` - Builds the app for production
- `npm run lint` - Runs ESLint
- `npm run format` - Formats code with Prettier
- `npm run analyze` - Analyzes bundle size

## 🧪 Testing

### Unit Tests
```bash
npm test
```

### Component Tests
```bash
npm run test:components
```

### E2E Tests
```bash
npm run test:e2e
```

## 🎨 UI Components

The application uses a combination of custom components and Material-UI:

- **Layout Components**
  - AppBar
  - Sidebar
  - Dashboard
  - Project Board

- **Task Components**
  - Task Card
  - Task List
  - Task Form
  - Task Details

- **Common Components**
  - Button
  - Input
  - Select
  - Modal
  - Loading
  - Error Boundary

## 🔐 Security

- CSRF protection
- XSS prevention
- Secure HTTP headers
- Input sanitization
- Token-based authentication
- Secure cookie handling

## 📦 Dependencies

Key dependencies:
- `@mui/material` - Material-UI components
- `@reduxjs/toolkit` - Redux state management
- `react-router-dom` - Routing
- `axios` - HTTP client
- `react-hook-form` - Form handling
- `date-fns` - Date manipulation
- `react-query` - Data fetching

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details.
