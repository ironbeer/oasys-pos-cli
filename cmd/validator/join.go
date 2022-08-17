package validator

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	"github.com/oasysgames/oasys-pos-cli/cmd/utils"
	"github.com/oasysgames/oasys-pos-cli/contracts"
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/oasysgames/oasys-pos-cli/util"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   cmdPrefix + "join",
	Short: "Join as a validator in the proof-of-stake.",
	Run: func(cmd *cobra.Command, args []string) {
		operator, err := cmd.Flags().GetString(constants.OperatorFlag)
		if err != nil {
			util.Fatal(err)
		}

		wallet, err := utils.NewWallet(cmd)
		if err != nil {
			util.Fatal(err)
		}

		doJoin(wallet, operator)
	},
}

func doJoin(wallet *eth.Wallet, operator string) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	txOpts, err := wallet.GetTransactOpts(ctx)
	if err != nil {
		util.Fatal(err)
	}

	stakemanager, err := contracts.NewStakeManager(wallet.Client)
	if err != nil {
		util.Fatal(err)
	}

	result, err := stakemanager.GetValidatorInfo(wallet.GetCallOpts(ctx), wallet.From)
	if err != nil {
		util.Fatal(err)
	}

	to := common.HexToAddress(operator)
	if result.Operator == to {
		util.Fatal(errors.New("already joined"))
	}

	tx, err := stakemanager.JoinValidator(txOpts, to)
	if err != nil {
		util.Fatal(err)
	}

	fmt.Printf("sending (tx: %s)...", tx.Hash().String())

	receipt, err := wallet.WaitForTransactionReceipt(ctx, tx.Hash())
	if err != nil {
		util.Fatal(err)
	}

	fmt.Printf(": success with %d gas\n", receipt.GasUsed)
}
