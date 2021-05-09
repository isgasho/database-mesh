# Database Mesh

[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![GoDoc](https://img.shields.io/badge/Godoc-reference-blue.svg)](https://godoc.org/github.com/SphereEx/database-mesh)
[![Go Report Card](https://goreportcard.com/badge/github.com/SphereEx/database-mesh)](https://goreportcard.com/report/github.com/SphereEx/database-mesh)
[![GitHub release](https://img.shields.io/github/tag/SphereEx/database-mesh.svg?label=release)](https://github.com/SphereEx/database-mesh/releases)
[![GitHub release date](https://img.shields.io/github/release-date/SphereEx/database-mesh.svg)](https://github.com/SphereEx/database-mesh/releases)
[![Coverage Status](https://codecov.io/gh/SphereEx/database-mesh/branch/master/graph/badge.svg)](https://codecov.io/gh/SphereEx/database-mesh)


## What is Database Mesh ?
Database Mesh defines itself as a cloud native database agent of the Kubernetes environment, in charge of all the access to the database in the form of sidecar. It provides a mesh layer interacting with the database, so consider this as Database Mesh.

Database Mesh emphasizes on how to connect distributed data-access applications with the databases. Focusing on interaction, it effectively organizes the interaction between messy applications and the databases. The applications and databases those use Database Mesh to visit databases will form a large grid system, where they just need to be put into the right positions accordingly. They are all governed by the mesh layer.
