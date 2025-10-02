package main

import (
    "context"
    "os"
    "os/signal"
    "syscall"
    "fmt"

	"github.com/Kirill-Sidorov94/wbl1/utils/sort/config"
    sorterService "github.com/Kirill-Sidorov94/wbl1/utils/sort/sorter"
    "github.com/Kirill-Sidorov94/wbl1/utils/sort/input"
    "github.com/Kirill-Sidorov94/wbl1/utils/sort/output"
)

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(1)
    }
}

func run() error {
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    defer stop()

    config, err := config.New()
    if err != nil {
        return err
    }

    errChan := make(chan error, 1)
    input := input.New(config.File, errChan)
    sorter := sorterService.New(&sorterService.Options{
        LinesChan: input.GetLinesChan(), 
        ErrChan: errChan, 
        CheckSorted: config.CheckSorted,
        Reverse: config.Reverse,
        Unique: config.Unique,
        SortType: config.SortType,
        KeyColumn: config.KeyColumn,
        IgnoreBlanks: config.IgnoreBlanks,
    })
    output := output.New(sorter.GetResultChan(), errChan)
    input.Read(ctx)
    sorter.Process(ctx)
    output.Write(ctx)

    select {
    case err := <-errChan:
        return fmt.Errorf("pipeline error: %w", err)
    case <-ctx.Done():
        return fmt.Errorf("interrupted: %w", ctx.Err())
    case <-output.Done():
    }

    return nil
}
