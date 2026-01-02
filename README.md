# ShellLog

ShellLog is a tool that logs **all shell input and output** to a file.

## Purpose

ShellLog is intended for **penetration testing engagements** where accurate evidence and reproducibility are required. 
It allows testers to reconstruct actions, justify findings, and reproduce results when documentation is incomplete or questions arise later.

## Important Warning

ShellLog logs **everything**, including **sensitive information** such as passwords, tokens, and keys.
All data is stored **in plain text**. This is **intentional**. You are fully responsible for handling and cleaning up the recorded data.
