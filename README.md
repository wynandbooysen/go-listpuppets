## go-listpuppets

[![Go Report Card](https://goreportcard.com/badge/github.com/wynandbooysen/go-listpuppets)](https://goreportcard.com/report/github.com/wynandbooysen/go-listpuppets)

Simple wrapper to expose command line 'puppetserver ca list --all' via API to get list of servers registered on Puppet Master

Built this due to Talend's default component for connecting via SSH (tSSH) still uses ganymed-ssh-2 (deprecated) which can't connect to servers after SSH hardening - turning off old insecure algorithms.

Quick and dirty way to get a list of all agents registered on the puppet master in JSON format for me to consume in Talend
