@echo off
echo 🌍 Assignment Toolkit Cross-Platform Setup
echo ==========================================

echo 📋 Detected: Windows 11
echo.

REM Check if Windows build exists
if not exist "builds\windows\assignment-toolkit.exe" (
    echo ❌ Windows build not found: builds\windows\assignment-toolkit.exe
    echo 🔧 Please build for Windows first:
    echo    go build -o builds\windows\assignment-toolkit.exe
    pause
    exit /b 1
)

REM Copy to current directory
copy "builds\windows\assignment-toolkit.exe" "assignment-toolkit.exe" >nul

echo ✅ Assignment Toolkit ready for Windows!
echo.
echo 🚀 Quick Test:
echo   assignment-toolkit.exe --help
echo.
echo 🎯 Next Steps:
echo   1. Create a workspace: mkdir C:\my-templates ^&^& cd C:\my-templates
echo   2. Initialize: C:\path\to\assignment-toolkit.exe init
echo   3. Create template: C:\path\to\assignment-toolkit.exe create essay
echo.
echo 📚 Documentation:
echo   docs\GETTING_STARTED_CHECKLIST.md
echo   docs\CROSS_PLATFORM_COMPATIBILITY_GUIDE.md
echo.
pause
