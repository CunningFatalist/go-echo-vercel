# Deploying a Go Echo application to Vercel

This is a simple example on how to deploy a Go Echo application to Vercel.

## Prerequisites

- [Vercel CLI](https://vercel.com/docs/cli)
- [Go](https://golang.org/dl/)

## Steps

### 1. Test your application locally

```bash
go run main.go
```

### 2. Deploy to Vercel

```bash
vercel

# or if you want to deploy to production
vercel --prod
```

### 3. Access your application

Just open the URL provided by Vercel after the deployment.

## Additional Resources

- [Vercel Docs](https://vercel.com/docs)
- [Vercel Go Runtime Docs](https://vercel.com/docs/functions/runtimes/go)
- [Go Echo](https://echo.labstack.com/)
- [Nice guide on Vercel and Go microservices](https://sorcererxw.com/en/articles/vercel-go-microservice)