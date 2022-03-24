import { PlusOutlined, DeleteOutlined, EditOutlined, EyeOutlined } from '@ant-design/icons';
import { Button, message, Input, Drawer } from 'antd';
import React, { useState, useRef } from 'react';
import { useIntl, FormattedMessage } from 'umi';
import { PageContainer, FooterToolbar } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { ModalForm, ProFormText, ProFormTextArea } from '@ant-design/pro-form';
import ProDescriptions from '@ant-design/pro-descriptions';
import UpdateForm from './components/UpdateForm';
import {
  product,
  productDetail,
  updateProduct,
  createNewProduct,
  removeProduct,
} from '@/services/ant-design-pro/api';

/**
 * @en-US Add node
 * @zh-CN 添加节点
 * @param fields
 */

// const handleAdd = async (fields) => {
//   const hide = message.loading('正在添加');

//   try {
//     await addRule({ ...fields });
//     hide();
//     message.success('Added successfully');
//     return true;
//   } catch (error) {
//     hide();
//     message.error('Adding failed, please try again!');
//     return false;
//   }
// };
// /**
//  * @en-US Update node
//  * @zh-CN 更新节点
//  *
//  * @param fields
//  */

// const handleUpdate = async (fields) => {
//   const hide = message.loading('Configuring');

//   try {
//     await updateRule({
//       name: fields.name,
//       desc: fields.desc,
//       key: fields.key,
//     });
//     hide();
//     message.success('Configuration is successful');
//     return true;
//   } catch (error) {
//     hide();
//     message.error('Configuration failed, please try again!');
//     return false;
//   }
// };
// /**
//  *  Delete node
//  * @zh-CN 删除节点
//  *
//  * @param selectedRows
//  */

// const handleRemove = async (selectedRows) => {
//   const hide = message.loading('正在删除');
//   if (!selectedRows) return true;

//   try {
//     await removeRule({
//       key: selectedRows.map((row) => row.key),
//     });
//     hide();
//     message.success('Deleted successfully and will refresh soon');
//     return true;
//   } catch (error) {
//     hide();
//     message.error('Delete failed, please try again');
//     return false;
//   }
// };

const ProductTable = () => {
  /**
   * @en-US Pop-up window of new window
   * @zh-CN 新建窗口的弹窗
   *  */
  const [createModalVisible, handleModalVisible] = useState(false);
  /**
   * @en-US The pop-up window of the distribution update window
   * @zh-CN 分布更新窗口的弹窗
   * */

  const [updateModalVisible, handleUpdateModalVisible] = useState(false);
  const [showDetail, setShowDetail] = useState(false);
  const [modalType, setModalType] = useState();
  const actionRef = useRef();
  const [currentRow, setCurrentRow] = useState();
  const [selectedRowsState, setSelectedRows] = useState([]);
  /**
   * @en-US International configuration
   * @zh-CN 国际化配置
   * */

  const handleRemoveProduct = async (id) => {
    try {
      const response = await removeProduct(id);
      if (response.status === 'ok') {
        message.success('Deleted successfully');
        actionRef.current.reload();
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const handleUpdateProduct = async (value) => {
    const payload = {
      name: value.name,
      description: value.description,
    };

    try {
      const response = await updateProduct(value.id, payload);
      if (response.status === 'ok') {
        message.success('Update user successfully');
        handleModalVisible(false);
        if (actionRef.current) {
          actionRef.current.reload();
        }
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const handleProductDetail = async (id) => {
    try {
      const response = await productDetail(id);
      if (response.status === 'ok') {
        setShowDetail(true);
        setCurrentRow(response.data);
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const handleCreateProduct = async (value) => {
    try {
      const response = await createNewProduct(value);
      if (response.status === 'ok') {
        message.success('Product created successfully');
        handleModalVisible(false);
        if (actionRef.current) {
          actionRef.current.reload();
        }
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const intl = useIntl();
  const columns = [
    {
      title: 'Name',
      dataIndex: 'Name',
    },
    {
      title: 'Description',
      dataIndex: 'Description',
    },
    {
      title: 'Status',
      dataIndex: 'Status',
    },
    {
      title: 'Action',
      dataIndex: 'option',
      valueType: 'option',
      render: (_, rowData) => {
        return (
          <div style={{ display: 'flex' }}>
            <div style={{ marginRight: 5 }}>
              <Button onClick={() => handleRemoveProduct(rowData.id)}>
                <DeleteOutlined />
              </Button>
            </div>
            <div style={{ marginRight: 5 }}>
              <Button
                onClick={() => {
                  handleModalVisible(true);
                  setCurrentRow(rowData);
                  setModalType('edit');
                }}
              >
                <EditOutlined />
              </Button>
            </div>
            <div style={{ marginRight: 5 }}>
              <Button
                onClick={() => {
                  handleProductDetail(rowData.id);
                }}
              >
                <EyeOutlined />
              </Button>
            </div>
          </div>
        );
      },
    },
  ];
  return (
    <PageContainer>
      <ProTable
        headerTitle={intl.formatMessage({
          id: 'pages.productTable.title',
          defaultMessage: 'Product List',
        })}
        actionRef={actionRef}
        rowKey="key"
        search={{
          labelWidth: 120,
        }}
        toolBarRender={() => [
          <Button
            type="primary"
            key="primary"
            onClick={() => {
              setCurrentRow(undefined);
              setModalType('add');
              handleModalVisible(true);
            }}
          >
            <PlusOutlined /> <FormattedMessage id="pages.searchTable.new" defaultMessage="New" />
          </Button>,
        ]}
        request={product}
        columns={columns}
        rowSelection={{
          onChange: (_, selectedRows) => {
            setSelectedRows(selectedRows);
          },
        }}
      />
      {/* {selectedRowsState?.length > 0 && (
      <FooterToolbar
        extra={
          <div>
            <FormattedMessage id="pages.productTable.chosen" defaultMessage="Chosen" />{' '}
            <a
              style={{
                fontWeight: 600,
              }}
            >
              {selectedRowsState.length}
            </a>{' '}
            <FormattedMessage id="pages.productTable.item" defaultMessage="项" />
            &nbsp;&nbsp;
            <span>
              <FormattedMessage
                id="pages.productTable.totalServiceCalls"
                defaultMessage="Total number of service calls"
              />{' '}
              {selectedRowsState.reduce((pre, item) => pre + item.callNo, 0)}{' '}
              <FormattedMessage id="pages.productTable.tenThousand" defaultMessage="万" />
            </span>
          </div>
        }
      >
        <Button
          onClick={async () => {
            await handleRemove(selectedRowsState);
            setSelectedRows([]);
            actionRef.current?.reloadAndRest?.();
          }}
        >
          <FormattedMessage
            id="pages.productTable.batchDeletion"
            defaultMessage="Batch deletion"
          />
        </Button>
        <Button type="primary">
          <FormattedMessage
            id="pages.productTable.batchApproval"
            defaultMessage="Batch approval"
          />
        </Button>
      </FooterToolbar>
    )} */}
      <ModalForm
        initialValues={{
          name: currentRow?.name,
          description: currentRow?.description,
        }}
        title={intl.formatMessage({
          id: 'pages.productTable.createForm.newRule',
          defaultMessage: modalType === 'edit' ? 'Edit Product' : 'Create Product',
        })}
        width="400px"
        visible={createModalVisible}
        onVisibleChange={handleModalVisible}
        onFinish={(value) =>
          modalType === 'edit'
            ? handleUpdateProduct({ ...value, id: currentRow.id })
            : handleCreateProduct(value)
        }
      >
        <ProFormText width="md" name="name" placeholder="Name" label="Name" />
        <ProFormText width="md" name="description" placeholder="Description" label="Description" />
      </ModalForm>

      <Drawer
        width={600}
        visible={showDetail}
        onClose={() => {
          setCurrentRow(undefined);
          setShowDetail(false);
        }}
        closable={false}
      >
        {currentRow?.name && (
          <ProDescriptions
            column={2}
            title={currentRow?.name.toUpperCase()}
            request={async () => ({
              data: currentRow || {},
            })}
            params={{
              id: currentRow?.name,
            }}
            columns={columns}
          />
        )}
      </Drawer>
    </PageContainer>
  );
};

export default ProductTable;
