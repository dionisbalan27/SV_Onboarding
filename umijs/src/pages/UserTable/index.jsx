import { EyeOutlined } from '@ant-design/icons';
import { Button, message, Input, Drawer } from 'antd';
import { useState, useRef } from 'react';
import { useIntl, FormattedMessage } from 'umi';
import { PageContainer, FooterToolbar } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { ModalForm, ProFormSelect, ProFormText } from '@ant-design/pro-form';
import ProDescriptions from '@ant-design/pro-descriptions';
import { user, userDetail } from '@/services/ant-design-pro/api';
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

const UserTabel = () => {
  /**
   * @en-US Pop-up window of new window
   * @zh-CN 新建窗口的弹窗
   *  */
  const [createModalVisible, handleModalVisible] = useState(false);
  /**
   * @en-US The pop-up window of the distribution update window
   * @zh-CN 分布更新窗口的弹窗
   * */
  const [roleData, setRoleData] = useState([]);
  const [updateModalVisible, handleUpdateModalVisible] = useState(false);
  const [showDetail, setShowDetail] = useState(false);
  const actionRef = useRef();
  const [currentRow, setCurrentRow] = useState();
  const [selectedRowsState, setSelectedRows] = useState([]);
  /**
   * @en-US International configuration
   * @zh-CN 国际化配置
   * */
  const handleUserDetail = async (id) => {
    try {
      const response = await userDetail(id);
      console.log(response.data);
      if (response.status === 'ok') {
        setShowDetail(true);
        setCurrentRow(response.data);
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const intl = useIntl();
  const columns = [
    {
      title: 'Name',
      dataIndex: 'name',
      render: (dom, entity) => {
        return (
          <a
            onClick={() => {
              handleUserDetail(entity.id);
            }}
          >
            {dom}
          </a>
        );
      },
    },
    {
      title: 'Role',
      dataIndex: 'role.title',
      render: (_, row) => {
        return <p>{row.role.title}</p>;
      },
    },
    {
      title: 'Role Id',
      dataIndex: 'role.id',
      render: (_, row) => {
        return <p>{row.role.id}</p>;
      },
    },
    {
      title: 'Active',
      dataIndex: 'active',
      render: (_, row) => {
        return <p>{row.active ? 'Active' : 'Inactive'}</p>;
      },
    },
    {
      title: 'Action',
      dataIndex: 'option',
      valueType: 'option',
      render: (_, rowData) => {
        return (
          <div style={{ display: 'flex' }}>
            <div style={{ marginRight: 5 }}>
              <Button
                onClick={() => {
                  handleUserDetail(rowData.id);
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
          id: 'pages.userTable.title',
          defaultMessage: 'User List',
        })}
        actionRef={actionRef}
        rowKey="key"
        search={{
          labelWidth: 120,
        }}
        request={user}
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
              <FormattedMessage id="pages.userTable.chosen" defaultMessage="Chosen" />{' '}
              <a
                style={{
                  fontWeight: 600,
                }}
              >
                {selectedRowsState.length}
              </a>{' '}
              <FormattedMessage id="pages.userTable.item" defaultMessage="项" />
              &nbsp;&nbsp;
              <span>
                <FormattedMessage
                  id="pages.userTable.totalServiceCalls"
                  defaultMessage="Total number of service calls"
                />{' '}
                {selectedRowsState.reduce((pre, item) => pre + item.callNo, 0)}{' '}
                <FormattedMessage id="pages.userTable.tenThousand" defaultMessage="万" />
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
            <FormattedMessage id="pages.userTable.batchDeletion" defaultMessage="Batch deletion" />
          </Button>
          <Button type="primary">
            <FormattedMessage id="pages.userTable.batchApproval" defaultMessage="Batch approval" />
          </Button>
        </FooterToolbar>
      )} */}

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

export default UserTabel;
