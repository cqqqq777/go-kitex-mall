package main

import (
	"context"
	"errors"
	"time"

	"github.com/cqqqq777/go-kitex-mall/cmd/product/dao"
	"github.com/cqqqq777/go-kitex-mall/cmd/product/model"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
	product "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/product"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/cqqqq777/go-kitex-mall/shared/response"

	"github.com/bwmarrin/snowflake"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct {
	Dao *dao.Product
	MerchantManager
	Producer
	OperateManager
}

type MerchantManager interface {
	GetInfo(ctx context.Context, MerchantId int64) (*common.Merchant, error)
}

type Producer interface {
	Produce(images []*model.Image) error
}

type OperateManager interface {
	GetProductOperateInfo(ctx context.Context, uid int64, pid int64) (*common.ProductOperateInfo, error)
}

// PublishProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) PublishProduct(ctx context.Context, req *product.MallPublishProductRequest) (resp *product.MallPublishProductResponse, err error) {
	// TODO: Your code here...
	resp = new(product.MallPublishProductResponse)
	productBasic := new(model.ProductBasic)

	// get merchant info
	merchant, err := s.MerchantManager.GetInfo(ctx, req.MerchantId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetMerchantInfo)
		log.Zlogger.Errorf("get merchant info failed err:%s", err.Error())
		return resp, nil
	}

	// generate product id
	sf, err := snowflake.NewNode(consts.ProductSnowflakeNode)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGenerateProductId)
		log.Zlogger.Errorf("generate product id failed err:%s", err.Error())
		return resp, nil
	}

	// build product
	productBasic.Mid = merchant.Id
	productBasic.Id = sf.Generate().Int64()
	productBasic.Name = req.Name
	productBasic.Description = req.Description
	productBasic.CreateTime = time.Now()
	productBasic.UpdateTime = time.Now()
	productBasic.Price = req.Price
	productBasic.Stock = req.Stock
	for _, v := range req.Images {
		productBasic.Images = append(productBasic.Images, &model.Image{
			Id:   v.Id,
			Path: v.Path,
		})
	}

	// create product in mongodb
	err = s.Dao.CreateProduct(ctx, productBasic)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrCreateProduct)
		log.Zlogger.Errorf("create product failed err:%s", err.Error())
		return resp, nil
	}

	// publish message to nsq
	err = s.Producer.Produce(productBasic.Images)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrPublishMsgInNsq)
		log.Zlogger.Errorf("publish msg to nsq failed err:%s", err.Error())
		return resp, nil
	}

	// build response
	resp.CommonResp = response.NewCommonResp(nil)
	resp.ProductId = productBasic.Id
	return resp, nil
}

// DelProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DelProduct(ctx context.Context, req *product.MallDelProductRequest) (resp *product.MallDelProductResponse, err error) {
	// TODO: Your code here...
	resp = new(product.MallDelProductResponse)

	// del product
	err = s.Dao.DelProduct(ctx, req.ProductId, req.MerchantId)
	if err != nil {
		if errors.Is(err, dao.ErrNoPermission) {
			resp.CommonResp.Code = errz.CodeNoPermission
			resp.CommonResp.Msg = err.Error()
			return resp, nil
		} else if errors.Is(err, dao.ErrNoProduct) {
			resp.CommonResp = response.NewCommonResp(errz.ErrNoProduct)
			return resp, nil
		}
		resp.CommonResp = response.NewCommonResp(errz.ErrProductInternal)
		return resp, nil
	}

	// clear cache
	err = s.Dao.ClearProductCache(ctx, req.ProductId)
	if err != nil {
		log.Zlogger.Errorf("clear product cache failed err:%s", err.Error())
	}

	resp.CommonResp = response.NewCommonResp(nil)
	return resp, nil
}

// ProductList implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductList(ctx context.Context, req *product.MallProductListRequest) (resp *product.MallProductListResponse, err error) {
	// TODO: Your code here...
	resp = new(product.MallProductListResponse)

	// check list key
	if req.Sort == "" {
		req.Sort = "time"
	}
	list, num, err := s.Dao.ProductList(ctx, req.Sort, int64(req.Page), int64(req.PageSize))
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrProductInternal)
		log.Zlogger.Errorf("get product list failed err:%s", err.Error())
		return resp, nil
	}

	// build response
	resp.TotalNum = num
	for _, v := range list {
		var respProduct *common.Product
		respProduct.Id = v.Id
		respProduct.MId = v.Mid
		respProduct.Name = v.Name
		respProduct.Price = v.Price
		respProduct.Description = v.Description
		respProduct.Stock = v.Stock
		respProduct.Status = v.Status
		for _, i := range v.Images {
			var image *common.Image
			image.Id = i.Id
			image.Path = i.Path
			respProduct.Iamges = append(respProduct.Iamges, image)
		}
		resp.Products = append(resp.Products, respProduct)
	}
	resp.CommonResp = response.NewCommonResp(nil)
	return resp, nil
}

// ProductDetail implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductDetail(ctx context.Context, req *product.MallProductDetailRequest) (resp *product.MallProductDetailResponse, err error) {
	// TODO: Your code here...
	resp = new(product.MallProductDetailResponse)

	// get basic info
	info, err := s.Dao.GetProductInfo(ctx, req.ProductId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrProductInternal)
		return resp, nil
	}
	resp.Product.BasicInfo.Id = info.Id
	resp.Product.BasicInfo.MId = info.Mid
	resp.Product.BasicInfo.Name = info.Name
	resp.Product.BasicInfo.Price = info.Price
	resp.Product.BasicInfo.Stock = info.Stock
	resp.Product.BasicInfo.Description = info.Description
	resp.Product.CreateTime = info.CreateTime.UnixNano()
	resp.Product.UpdateTime = info.UpdateTime.UnixNano()
	for _, i := range info.Images {
		var image *common.Image
		image.Id = i.Id
		image.Path = i.Path
		resp.Product.BasicInfo.Iamges = append(resp.Product.BasicInfo.Iamges, image)
	}

	// get merchant info
	resp.Product.MerchantInfo, err = s.MerchantManager.GetInfo(ctx, info.Mid)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetMerchantInfo)
		log.Zlogger.Errorf("get merchant info failed err:%s", err.Error())
		return resp, nil
	}

	// get operate info
	resp.Product.OperateInfo, err = s.OperateManager.GetProductOperateInfo(ctx, req.UserId, req.ProductId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetOperateInfo)
		log.Zlogger.Errorf("get operate info failed err:%s", err.Error())
		return resp, nil
	}

	// cache product detail
	err = s.Dao.CacheProductInfo(ctx, info.Id, resp.Product)
	if err != nil {
		log.Zlogger.Errorf("cache product detail failed err:%s", err.Error())
	}

	resp.CommonResp = response.NewCommonResp(nil)
	return resp, nil
}

// SearchProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) SearchProduct(ctx context.Context, req *product.MallSearchProductRequest) (resp *product.MallSearchProductResponse, err error) {
	// TODO: Your code here...
	resp = new(product.MallSearchProductResponse)

	// search
	list, err := s.Dao.SearchProduct(ctx, req.Key)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrProductInternal)
		log.Zlogger.Errorf("search product failed err:%s", err.Error())
		return resp, nil
	}

	// build response
	for _, v := range list {
		var respProduct *common.Product
		respProduct.Id = v.Id
		respProduct.MId = v.Mid
		respProduct.Name = v.Name
		respProduct.Price = v.Price
		respProduct.Description = v.Description
		respProduct.Stock = v.Stock
		respProduct.Status = v.Status
		for _, i := range v.Images {
			var image *common.Image
			image.Id = i.Id
			image.Path = i.Path
			respProduct.Iamges = append(respProduct.Iamges, image)
		}
		resp.Products = append(resp.Products, respProduct)
	}
	resp.CommonResp = response.NewCommonResp(nil)

	return resp, nil
}

// ProductFavoriteList implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductFavoriteList(ctx context.Context, req *product.MallProductFavoriteListRequest) (resp *product.MallProductFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishedProducts implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) PublishedProducts(ctx context.Context, req *product.MallProductPublishedListRequest) (resp *product.MallProductPublishedListResponse, err error) {
	// TODO: Your code here...
	resp = new(product.MallProductPublishedListResponse)

	// get list
	list, err := s.Dao.PublishedProducts(ctx, req.MerchantId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrProductInternal)
		log.Zlogger.Errorf("get published products failed err:%s", err.Error())
		return resp, nil
	}

	// build response
	for _, v := range list {
		var respProduct *common.Product
		respProduct.Id = v.Id
		respProduct.MId = v.Mid
		respProduct.Name = v.Name
		respProduct.Price = v.Price
		respProduct.Description = v.Description
		respProduct.Stock = v.Stock
		respProduct.Status = v.Status
		for _, i := range v.Images {
			var image *common.Image
			image.Id = i.Id
			image.Path = i.Path
			respProduct.Iamges = append(respProduct.Iamges, image)
		}
		resp.Products = append(resp.Products, respProduct)
	}
	resp.CommonResp = response.NewCommonResp(nil)

	return resp, nil
}
