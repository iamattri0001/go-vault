const VaultsList = ({ vaults, setVaults }) => {
  return (
    <div className="flex-grow flex items-center justify-center">
      {vaults.length === 0 ? (
        <p className="text-center text-gray-500 ">
          No vaults available. Try creating a new vault.
        </p>
      ) : (
        <div>{/* Render vaults here */}</div>
      )}
    </div>
  );
};

export default VaultsList;
