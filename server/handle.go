package services

import (
	"context"
	"fmt"
	"github.com/L-DENG/rpcService/server/address"
	"strconv"

	"github.com/L-DENG/rpcService/protobuf/wallet"
)

func (s *RpcServer) GetSupportCoins(ctx context.Context, in *wallet.SupportCoinsRequest) (*wallet.SupportCoinsResponse, error) {

	return &wallet.SupportCoinsResponse{
		Code:    strconv.Itoa(200),
		Msg:     "success request",
		Support: true,
	}, nil
}
func (s *RpcServer) GetWalletAddress(ctx context.Context, in *wallet.WalletAddressRequest) (*wallet.WalletAddressResponse, error) {
	addressInfo, err := address.CreateAddressFromPrivateKey()
	if err != nil {
		fmt.Println("err create address")
		return &wallet.WalletAddressResponse{
			Code:      strconv.Itoa(400),
			Msg:       "create address fail",
			Address:   "",
			PublicKey: "",
		}, nil
	}
	return &wallet.WalletAddressResponse{
		Code:      strconv.Itoa(200),
		Msg:       "success request",
		Address:   addressInfo.Address,
		PublicKey: addressInfo.PublicKey,
	}, nil
}
